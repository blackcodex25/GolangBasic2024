package main

import (
	"fmt"
	"sync"
)

/*
	ในแบบฝึกหัดนี้ เราจะใช้คุณสมบัติการทำงานพร้อมกันของ Go เพื่อทำให้

เว็บครอว์เลอร์ทำงานแบบขนาน
ปรับเปลี่ยนฟังก์ชัน Crawl เพื่อดึงข้อมูลจาก URL หลายๆ แห่งพร้อมกัน
โดยไม่ดึง URL เดียวกันซ้ำสองครั้ง
คำแนะนำ: คุณสามารถเก็บแคชของ URL ที่ถูกดึงข้อมูลแล้วใน map แต่การใช้ map เพียง
อย่างเดียวไม่ปลอดภัยสำหรับการใช้งานพร้อมกัน
*/

// Fetcher interface สำหรับการดึงข้อมูลจาก URL
type Fetcher interface {
	// Fecth ส่งคืนเนื้อหาของ URL
	// และส่วนของ URL ที่พบในหน้านั้น
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCache ใช้เพื่อเก็บ URL ที่ถูกดึงข้อมูลแล้ว
type SafeCache struct {
	mutex sync.Mutex
	cache map[string]bool
}

// NewSafeCache สร้าง SafeCache ใหม่
func NewSafeCache() *SafeCache {
	return &SafeCache{
		cache: make(map[string]bool),
	}
}

// Visited ตรวจสอบว่า URL ถูกดึงข้อมูลแล้วหรือยัง และทำการล็อค
func (c *SafeCache) Visited(url string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cache[url] {
		return true
	}
	c.cache[url] = true
	return false
}

// การรวบรวมข้อมูลใช้ตัวดึงข้อมูลเพิ่อรวบรวมข้อมูลแบบเรียกซ้ำ
// หน้าเว็บที่เริ่มต้นด้วย URL จะถูกเข้าถึงได้สูงสุดตามความลึกที่กำหนด
// Crawl ดึงข้อมูลจาก URL หลายๆ แห่งพร้อมกันโดยไม่ดึง
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, wg *sync.WaitGroup) {
	// TODO: ดึงข้อมูลจาก URL หลายๆ แห่งพร้อมกัน
	// TODO: หลีกเลี่ยงการดึงข้อมูลจาก URL เดียวกันซ้ำสองครั้ง
	defer wg.Done()
	// การใช้งานในตอนนี้ยังไม่ได้ทำทั้งสองสิ่งนี้:
	// หยุดการทำงานถ้า depth <= 0 หรือ URL นี้ถูกดึงข้อมูลแล้ว
	if depth <= 0 || cache.Visited(url) {
		return
	}

	// ดึงข้อมูลจาก URL
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("not found:", url)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// สร้าง goroutine ใหม่สำหรับการดึงข้อมูล URL ต่อไป
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, cache, wg)
	}
}

// ใช้ WaitGroup เพื่อติดตามการทำงานทั้งหมด
var wg sync.WaitGroup

func main() {
	// สร้าง SafeCache เพื่อป้องกันการดึง URL ซ้ำ
	cache := NewSafeCache()

	// เริ่มการดึงข้อมููลที่ URL แรก
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, cache, &wg)

	// รอการทำงานทั้งหมดให้เสร็จสิ้น
	wg.Wait()
}

// fakeFetcher เป็น Fetcher ที่คืนค่าผลลัพธ์ที่เตรียมไว้ล่วงหน้า
type fakeFetcher map[string]*fakeResult

// โครงสร้าง fakeResult เป็นโครงสร้างข้อมูล (struct) ที่มีสองฟิลด์
type fakeResult struct {
	body string   // ฟิลด์ body ชนิด string เก็บเนื้อหาของหน้าเว็บที่ถูกดึงข้อมูลจาก URL
	urls []string // ฟิลด์ urls ชนิด string เก็บรายการ URL ที่พบในหน้าเว็บนั้นๆ
}

// ฟังก์ชัน Fecth ทำหน้าที่ดึงข้อมูลจาก URL โดยรับค่า url ชนิด string ซึ่งเป็น URL ที่ต้องการดึงข้อมูล
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		// คืนค่ากลับเป็น body (เนื้อหาเว็บ), urls (รายการ URL ในหน้าเว็บ),
		// และ error หากเกิดข้อผิดพลาด
		return res.body, res.urls, nil
	}
	// แสดงข้อความเมื่อไม่พบ URL ใน map
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetch คือ fakeFetcher ที่มีข้อมูลอยู่แล้ว
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
