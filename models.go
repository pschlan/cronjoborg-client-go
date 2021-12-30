package cronjoborg

type JobSchedule struct {
	Timezone string `json:"timezone,omitempty"`
	Hours    []int  `json:"hours,omitempty"`
	MDays    []int  `json:"mdays,omsitempty"`
	Minutes  []int  `json:"minute,omitempty"`
	Months   []int  `json:"months,omitempty"`
	WDays    []int  `json:"wdays,omitempty"`
}

type JobAuth struct {
	Enable   bool   `json:"enable,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type JobNotificationSettings struct {
	OnFailure bool `json:"onFailure,omitempty"`
	OnSuccess bool `json:"onSuccess,omitempty"`
	OnDisable bool `json:"onDisable,omitempty"`
}

type JobExtendedData struct {
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}

const (
	JOBSTATUS_UNKNOWN         int = 0
	JOBSTATUS_OK              int = 1
	JOBSTATUS_FAILED_DNS      int = 2
	JOBSTATUS_FAILED_CONNECT  int = 3
	JOBSTATUS_FAILED_HTTP     int = 4
	JOBSTATUS_FAILED_TIMEOUT  int = 5
	JOBSTATUS_FAILED_SIZE     int = 6
	JOBSTATUS_FAILED_URL      int = 7
	JOBSTATUS_FAILED_INTERNAL int = 8
	JOBSTATUS_FAILED_UNKNOWN  int = 9
)

const (
	JOBTYPE_DEFAULT int = 0
	JOBTYPE_MONITOR int = 1
)

const (
	REQUESTMETHOD_GET     int = 0
	REQUESTMETHOD_POST    int = 1
	REQUESTMETHOD_OPTIONS int = 2
	REQUESTMETHOD_HEAD    int = 3
	REQUESTMETHOD_PUT     int = 4
	REQUESTMETHOD_DELETE  int = 5
	REQUESTMETHOD_TRACE   int = 6
	REQUESTMETHOD_CONNECT int = 7
	REQUESTMETHOD_PATCH   int = 8
)

type Job struct {
	JobID          int                     `json:"jobId,omitempty"`
	Enabled        bool                    `json:"enabled,omitempty"`
	Title          string                  `json:"title,omitempty"`
	SaveResponses  bool                    `json:"saveResponses,omitempty"`
	URL            string                  `json:"url,omitempty"`
	LastStatus     int                     `json:"lastStatus,omitempty"`
	LastDuration   int                     `json:"lastDuration,omitempty"`
	LastExecution  int                     `json:"lastExecution,omitempty"`
	NextExecution  int                     `json:"nextExecution,omitempty"`
	Type           int                     `json:"type,omitempty"`
	RequestTimeout int                     `json:"requestTimeout,omitempty"`
	Schedule       JobSchedule             `json:"schedule,omitempty"`
	RequestMethod  int                     `json:"requestMethod,omitempty"`
	Auth           JobAuth                 `json:"auth,omitempty"`
	Notification   JobNotificationSettings `json:"notification,omitempty"`
	ExtendedData   JobExtendedData         `json:"extendedData,omitempty"`
}
