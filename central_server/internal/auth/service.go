package auth

type Service interface {
	Login(authUser Peer) (string, error)
	Logout(authUser PeerLogOut) error
	GetUser(username string) (Peer, error)
	AssignPeer(string) (location string)
}

type ServiceClient struct {
	repository RepositoryAuth
	peerCount int
}

func NewServiceClient(repo RepositoryAuth) *ServiceClient {
	return &ServiceClient{repository: repo}

}
func (s ServiceClient) Login(authUser Peer) (string, error) {

	return s.repository.SavePeer(authUser)
}

func (s ServiceClient) Logout(authUser PeerLogOut) error {
	currentUser, err := s.repository.GetPeer(authUser.Username)
	if err != nil {
		return err
	}
	currentUser.State = "down"
	err = s.repository.UpdatePeer(currentUser)
	return err
}

func (s ServiceClient) GetUser(username string) (Peer, error) {
	return s.repository.GetPeer(username)
}
func (s ServiceClient) AssignPeer(string) (location string){
	list:= s.repository.PeerOrderList()
	if len(list) == 0 {
        return ""
    }
	c:=s.peerCount%len(list)
	candidatePeerUsername:= list[c]
	peer,err:=s.GetUser(candidatePeerUsername)
	if err == nil && peer.State == "up"{
		s.peerCount++
		return peer.UserURL
	}
	
	for peer.State == "down" && c<len(list){
		c++
		peer,_=s.GetUser(list[c])
		s.peerCount++
	}
	
	return 
	
	
}
