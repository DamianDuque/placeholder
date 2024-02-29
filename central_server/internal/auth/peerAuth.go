package auth

type Peer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserURL  string `json:"user_url"`
	State    string `json:"-"`
}
type PeerLogOut struct {
	Username string `json:"username"`
}
