package existence_conf

//import (
//	dpfm_api_input_reader "data-platform-api-orders-updates-rmq-kube/DPFM_API_Input_Reader"
//	"strconv"
//	"sync"
//
//	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
//	"golang.org/x/xerrors"
//)
//
//func (c *ExistenceConf) itemProjectExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
//	defer wg.Done()
//	wg2 := sync.WaitGroup{}
//	exReqTimes := 0
//
//	items := input.Header.Item
//	for _, item := range items {
//		project := getItemProjectExistenceConfKey(mapper, &item, exconfErrMsg)
//		wg2.Add(1)
//		exReqTimes++
//		go func() {
//			if isZero(project) {
//				wg2.Done()
//				return
//			}
//			res, err := c.projectExistenceConfRequest(project, mapper, input, existenceMap, mtx, log)
//			if err != nil {
//				mtx.Lock()
//				*errs = append(*errs, err)
//				mtx.Unlock()
//			}
//			if res != "" {
//				*exconfErrMsg = res
//			}
//			wg2.Done()
//		}()
//	}
//	wg2.Wait()
//	if exReqTimes == 0 {
//		*existenceMap = append(*existenceMap, false)
//	}
//}
//
//func (c *ExistenceConf) projectExistenceConfRequest(project string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
//	keys := newResult(map[string]interface{}{
//		"Project": project,
//	})
//	exist := false
//	defer func() {
//		mtx.Lock()
//		*existenceMap = append(*existenceMap, exist)
//		mtx.Unlock()
//	}()
//
//	req, err := jsonTypeConversion[Returns](input)
//	if err != nil {
//		return "", xerrors.Errorf("request create error: %w", err)
//	}
//	req.ProjectReturn.Project = project
//
//	exist, err = c.exconfRequest(req, mapper, log)
//	if err != nil {
//		return "", err
//	}
//	if !exist {
//		return keys.fail(), nil
//	}
//	return "", nil
//}
//
//func getItemProjectExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) string {
//	var project int
//
//	switch mapper.Field {
//	//case "Project":
//	//	if item.Project == nil {
//	//		project = 0
//	//	}
//	//	project = *item.Project
//	}
//	return strconv.Itoa(project)
//}
