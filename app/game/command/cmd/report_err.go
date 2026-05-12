package cmd

type ReportErr struct {
	Err error
}

func NewReportErr(err error) *ReportErr {
	return &ReportErr{Err: err}
}

func (r *ReportErr) Kind() string {
	return "ReportErr"
}

func (r *ReportErr) Error() string {
	return r.Err.Error()
}
