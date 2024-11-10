package main

import (
    "testing"
    "strings"
)

func TestMainFunction(t *testing.T) {
    // 테스트 목적으로 출력 캡처
    var output strings.Builder
    oldStdout := stdOut
    stdOut = &output
    defer func() { stdOut = oldStdout }()

    // main 함수 실행
    main()

    // 테스트 검증
    expectedOutputs := []string{
        "You are an adult. Access granted.",
        "Standard user permissions apply.",
    }

    for _, expected := range expectedOutputs {
        if !strings.Contains(output.String(), expected) {
            t.Errorf("Expected output %q not found", expected)
        }
    }
}
