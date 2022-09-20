Name:           gopic
Version:        0.0.1
Release:        1%{?dist}
Summary:        An opensource picPool tool
BuildArch:      %{_arch}

License:        MIT
Source0:        %{name}-%{version}.tar.gz

Requires:       bash

%description
RPM build for gopic

%install
install -Dpm0755 ../../out/gopic-linux-amd64 $RPM_BUILD_ROOT/%{_bindir}/gopic

%clean
rm -rf $RPM_BUILD_ROOT

%files
%{_bindir}/gopic

%changelog
* Tue Sep 20 2022 Rickylss <xiaohaibiao331@outlook.com> - 0.0.1
- First version being packaged