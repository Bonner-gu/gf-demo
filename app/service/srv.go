package service

var (
	AuthSrvCli    *AuthSrv
	IdmakerSrvCli *IdmakerSrv
	AdminSrvCli   *AdminSrv
)

func init() {
	AuthSrvCli = NewAuthSrv()
	IdmakerSrvCli = NewIdmakerSrv()
	AdminSrvCli = NewAdminSrv()
}
