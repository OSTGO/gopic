package plugin

import (
	"github.com/hirochachacha/go-smb2"
	"gopic/conf"
	"gopic/utils"
	"net"
	"os"
	"path"
)

type SambaStorage struct {
	B *utils.BaseStorage
}

const (
	sambaPluginName = "samba"
)

var sambaConfig map[string]interface{}

func (g *SambaStorage) Upload(im *utils.Image) (string, error) {
	responseURL := sambaConfig["responseurl"].(string)
	userName := sambaConfig["user"].(string)
	password := sambaConfig["password"].(string)
	picDir := sambaConfig["dir"].(string)
	address := sambaConfig["address"].(string)
	shareName := sambaConfig["share"].(string)
	return responseURL + path.Join(picDir, im.OutSuffix), uploadPictureToSamba(address, userName, shareName, password, picDir, im.OutSuffix, im.OutBytes)
}

func NewSambaStorage() *SambaStorage {
	return &SambaStorage{utils.NewBaseStorage()}
}

func uploadPictureToSamba(address, userName, shareName, password, picDir, suffix string, data []byte) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     userName,
			Password: password,
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		return err
	}
	defer func(s *smb2.Session) {
		_ = s.Logoff()
	}(s)

	fs, err := s.Mount(shareName)
	if err != nil {
		return err
	}
	defer func(fs *smb2.Share) {
		_ = fs.Umount()
	}(fs)
	_picPath := path.Join(picDir, suffix)
	createDirIfNotExist(fs, path.Dir(_picPath))

	f, err := fs.Create(_picPath)
	if err != nil {
		return err
	}
	defer func(f *smb2.File) {
		_ = f.Close()
	}(f)

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func createDirIfNotExist(fs *smb2.Share, p string) {
	createDir := func(pa string) {
		dd, err := fs.Open(pa)
		if err != nil {
			_ = fs.Mkdir(pa, os.ModeDir)
		} else {
			_ = dd.Close()
		}
	}
	for p != "." {
		defer createDir(p)
		p = path.Dir(p)
	}
}

func init() {
	utils.StroageHelp[sambaPluginName] = sambaHelp()
	sambaConfig = conf.Viper.GetStringMap(sambaPluginName)
	if sambaConfig == nil {
		return
	}
	active := sambaConfig["active"]
	if active == nil {
		return
	}
	if active == true {
		utils.StroageMap[sambaPluginName] = NewSambaStorage()
	}
}

func sambaHelp() string {
	return "samba plugin need this parameters:\nactive: false or true\nresponseURL: like \"http://192.168.200.100/mirror/\nuser: like smb\npassword: like password\nshare: like share\ndir: like my/test\naddress: like 192.168.200.100:445"
}
