package godo

import (
	"github.com/compermane/ic-go/pkg/domain/executor"
)

/*  Wraps godo execution algorithm inside this package.
 *	:param fns: Functions to execute
 *	:param rcvs: Receivers associated with them functions
 *	:param algorithm: Algorithm used for function execution (b1, fd)
 *  :param no_runs: Limit of iterations
 *  :param timeout: Time limit for iterations. Will not be utilized if no_runs is not 0
 */
func GODO(fns, rcvs []any, algorithm string, no_runs, timeout int) {
	executor.ExecuteFuncs(fns, rcvs, algorithm, no_runs, timeout)
}