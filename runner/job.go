package runner

// JobRequest models a request to run a job
type JobRequest struct {
	Image string
	Cmd   []string
	Stdin []byte
	Env   EnvVars
}

// JobResponse models a request to run a job
type JobResponse struct {
	Stdout []byte
}

// Execute runs the job
func (req *JobRequest) Execute(engine containerEngine) (*JobResponse, error) {
	container, err := engine.Run(req)
	if err != nil {
		return nil, err
	}

	return &JobResponse{
		Stdout: container.Stdout(),
	}, nil
}
