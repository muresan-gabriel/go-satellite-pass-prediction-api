package pass

type Service interface {
	Compute(req Request) ([]Pass, error)
}

type StubService struct{}

func NewStubService() *StubService {
	return &StubService{}
}

func (s *StubService) Compute(req Request) ([]Pass, error) {
	return []Pass{
		{
			AOS:          req.Start,
			LOS:          req.End,
			MaxElevation: 42.5,
			DurationSec:  480,
			Visibility:   "NIGHT",
		},
	}, nil
}
