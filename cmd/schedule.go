package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"

	"github.com/CodeSingerGnC/EbbinghausHelper/common"
)

type FileDataMap struct {
	Map map[string]Data
}

func NewFileDataMap() *FileDataMap {
	return &FileDataMap{
		Map: make(map[string]Data),
	}
}

func (m *FileDataMap) schedule() {
	m.getData()
	for filename, value := range m.Map {
		frequency := 10
		if value.Frequency > 0 {
			frequency = value.Frequency
		} 
		if !value.Initialized {
			count := 0
			addDay := 0
			for i := range value.Items {
				if count >= frequency {
					count = 0
					addDay += 1
				}
				value.Items[i].Scheduled = time.Now().AddDate(0, 0, addDay)
				value.Items[i].ThisTime = value.Items[i].Scheduled
				count += 1
			}
			value.Initialized = true
			stream, err := value.Marshal()
			if err != nil {
				return
			}
			m.write(filename, stream)
		}
		now := time.Now()
		for i := range value.Items {
			if value.Items[i].Scheduled.Before(now) {
				t := value.Items[i].Times
				value.Items[i].ThisTime = value.Items[i].Scheduled
				if t < len(value.ReviewScheme) {
					value.Items[i].Scheduled = time.Now().AddDate(0, 0, value.ReviewScheme[0])
				} else {
					value.Items[i].Scheduled = time.Now().AddDate(100, 0, 0)
				}
				value.Items[i].Times += 1
			}
		}
		stream, err := value.Marshal()
		if err != nil {
			return
		}
		m.write(filename, stream)
	}
}

func (m *FileDataMap) show() {
	fmt.Printf("Time: %s\n", time.Now().Format("2006-01-02 15:04"))
	m.getData()
	for filename, value := range m.Map {
		data := make([][]string, len(value.Items))
		for _, v := range value.Items {
			if EqualDate(time.Now(), v.ThisTime) {
				data = append(data, []string{v.Name, v.ThisTime.Format("2006-01-02"), v.Scheduled.Format("2006-01-02"), v.Website, fmt.Sprint(v.Times), v.Extra})
			}
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "ThisTime", "NextTime", "Web", "Times", "Extra"})
		for _, v := range data {
			table.Append(v)
		}
		filename = strings.ReplaceAll(filename, ".json", "")
		fmt.Printf("Schedule Table: %s\n", filename)
		table.Render()
	}
}

func (m *FileDataMap) getData() {
	if m == nil {
		m = &FileDataMap{
			Map: make(map[string]Data),
		}
	}

	dir, _ := common.GetBaseDir()
	jsonFiles, err := common.GetJsonFiles(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, jsonFile := range jsonFiles {
		f, err := os.Open(filepath.Join(dir, jsonFile.Name()))
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		content, err := io.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := Data{}

		data.Unmarshal(string(content))

		if m.Map == nil {
			m.Map = make(map[string]Data)
		}

		m.Map[jsonFile.Name()] = data
	}
}

func (m *FileDataMap) write(filename string, stream string) error {
	baseDir, err := common.GetBaseDir()
	if err != nil {
		return err
	}
	path := filepath.Join(baseDir, filename)
	file, _ := os.Create(path)
	file.WriteString(stream)
	file.Sync()
	return nil
}

func EqualDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}