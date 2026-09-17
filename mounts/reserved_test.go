package mounts

import "testing"

func TestValidateContainerMountTarget(t *testing.T) {
	cases := []struct {
		target string
		ok     bool
	}{
		{"/workspace/src", true},
		{"/data", true},
		{"/osg", false},
		{"/osg/data", false},
		{"/osg/policy.yaml", false},
		{"/proc", false},
		{"/sys/fs", false},
		{"/dev/null", false},
		{"/run/osg/ssh.sock", false},
		{"relative", false},
		{"/workspace/../osg", false}, // has ..
	}
	for _, tc := range cases {
		err := ValidateContainerMountTarget(tc.target)
		if tc.ok && err != nil {
			t.Fatalf("%s: unexpected err %v", tc.target, err)
		}
		if !tc.ok && err == nil {
			t.Fatalf("%s: expected error", tc.target)
		}
	}
}

func TestValidateUploadDest(t *testing.T) {
	if err := ValidateUploadDest("/workspace/file.txt"); err != nil {
		t.Fatal(err)
	}
	if err := ValidateUploadDest("/workspace"); err != nil {
		t.Fatal(err)
	}
	if err := ValidateUploadDest("/osg/data/x"); err == nil {
		t.Fatal("expected refuse /osg")
	}
}

func TestPathsOverlap(t *testing.T) {
	if !PathsOverlap("/osg", "/osg/data") {
		t.Fatal("expected overlap")
	}
	if PathsOverlap("/workspace", "/osg") {
		t.Fatal("no overlap")
	}
}
