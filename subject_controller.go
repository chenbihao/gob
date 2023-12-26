package main

import "gob/framework"

func SubjectAddController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectNameController")
	return nil
}
