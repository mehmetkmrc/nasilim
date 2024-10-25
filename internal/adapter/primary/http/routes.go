package http



func (s *server) SetupRoutes(){

}

func (s *server) webSetUp(){
	route := s.app.Group("/web")
	route.Get("/signup", s.security)
}