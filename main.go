package main

import (
	dpfm_api_caller "data-platform-api-delivery-document-csv-converter-rmq-kube/DPFM_API_Caller"
	dpfm_api_input_reader "data-platform-api-delivery-document-csv-converter-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-delivery-document-csv-converter-rmq-kube/config"
	"fmt"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()

	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.SessionControlQueue(), conf.RMQ.QueueToSQL(), 0)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	caller := dpfm_api_caller.NewDPFMAPICaller(conf, l)

	reader := dpfm_api_input_reader.NewFileReader()
	l.Info("READY!")
	for msg := range iter {
		l.Info("request is : %s\n", msg.Raw())
		start := time.Now()
		msg.Success()
		f, err := reader.ReadInput(msg.Raw())
		if err != nil {
			l.Error(err)
			continue
		}

		err = callProcess(rmq, caller, reader, f)
		if err != nil {
			l.Error(err)
			continue
		}

		l.Info("process time %v\n", time.Since(start).Milliseconds())
	}
}

func recovery(l *logger.Logger, err *error) {
	if err != nil {
		if *err != nil {
			l.Error("%+v", *err)
			if e := recover(); e != nil {
				*err = fmt.Errorf("error occurred: %w", e)
				l.Error(err)
				return
			}
		}
	}
}
func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func callProcess(rmq *rabbitmq.RabbitmqClient, caller *dpfm_api_caller.DPFMAPICaller, reader *dpfm_api_input_reader.FileReader, f *dpfm_api_input_reader.Request) error {
	var err error

	switch f.ServiceLabel {
	case "FUNCTION_DPFM_DELIVERY_DOCUMENT_CSV_READS":
		err = caller.DeliveryDocument(rmq, reader, f)
	default:
		return xerrors.Errorf("unknown API %s", f.ServiceLabel)
	}

	return err
}
