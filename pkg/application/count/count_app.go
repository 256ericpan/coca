package count

import (
	"github.com/phodal/coca/pkg/domain/jdomain"
)

func BuildCallMap(parserDeps []jdomain.JClassNode) map[string]int {
	var projectMethods = make(map[string]string)
	for _, clz := range parserDeps {
		clz.BuildStringMethodMap(projectMethods)
	}

	// TODO: support identify data class
	var callMap = make(map[string]int)
	for _, clz := range parserDeps {
		for _, method := range clz.Methods {
			for _, call := range method.MethodCalls {
				callMethod := call.BuildFullMethodName()
				if _, ok := projectMethods[callMethod]; ok {
					if callMap[callMethod] == 0 {
						callMap[callMethod] = 1
					} else {
						callMap[callMethod]++
					}
				}
			}
		}
	}

	return callMap
}

