/*package main

import "github.com/Didafe/crud-go/crud-go/src/configuration/rest_err"

func test(message string) (err *rest_err.RestErr) {

	err.Code

}*/

package controller

import "github.com/Didafe/crud-go/crud-go/src/configuration/rest_err"

func test(message string) (err *rest_err.RestErr) {
	if err != nil {
		// Se `err` não for nulo, então você pode acessar a propriedade `Code`
		_ = err.Code
	}
	return err
}

