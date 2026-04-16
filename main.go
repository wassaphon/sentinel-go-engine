package main

import (
	"fmt"
	"net/http"
	"time"
)

// ฟังก์ชันสำหรับเช็คสถานะเว็บไซต์ (Status Checker)
func checkStatus(url string) {
	// ตั้งค่าเวลา Timeout 5 วินาที ถ้าเกินนี้ถือว่าเว็บมีปัญหา
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("[ALERT] ❌ %s: เข้าถึงไม่ได้! (Error: %v)\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("[SAFE] ✅ %s: สถานะ %d (ใช้เวลา %v)\n", url, resp.StatusCode, duration)
}

func main() {
	// รายชื่อเว็บไซต์หรือระบบในเครือข่าย สพป. ที่คุณต้องการเฝ้าระวัง
	// คุณสามารถเพิ่ม url ของเว็บ pr.utt1.go.th หรือเว็บอื่นเข้าไปได้เลย
	targets := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://wordpress.org",
	}

	fmt.Println("--- SENTINEL Go Engine: Starting System Monitoring ---")
	fmt.Printf("ตรวจสอบทั้งหมด %d รายการ...\n\n", len(targets))

	// ใช้ 'go' keyword เพื่อเริ่มทำงานแบบขนาน (Goroutines)
	for _, url := range targets {
		go checkStatus(url)
	}

	// รอให้ระบบทำงานเสร็จ (หน่วงเวลาไว้ครู่หนึ่งเพื่อให้ Go เช็คครบทุกเว็บ)
	time.Sleep(10 * time.Second)
	fmt.Println("\n--- การสแกนเสร็จสิ้น ---")
}
