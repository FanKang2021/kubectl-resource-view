package writer

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

//NodeWrite
func NodeWrite(data [][]string, resourceType []string, outType bool) {
	//var table *tablewriter.Table

	table := table(outType)
	var header []string
	header = append(header, "NODE")
	
	// 如果没有指定资源类型，使用默认的全部类型
	if len(resourceType) == 0 {
		resourceType = []string{"cpu", "memory", "gpu", "pod"}
	}
	
	for _, t := range resourceType {
		switch {
		case t == "cpu":
			header = append(header,
				"CPU USE", "CPU REQ", "CPU REQ(%)", "CPU LIM", "CPU LIM(%)",
			)
		case t == "memory":
			header = append(header,
				"MEM USE", "MEM REQ", "MEM REQ(%)", "MEM LIM", "MEM LIM(%)",
			)
		case t == "gpu":
			header = append(header,
				"NVIDIA/GPU REQ", "NVIDIA/GPU REQ(%)", "NVIDIA/GPU LIM", "NVIDIA/GPU LIM(%)", "GPU MODEL",
			)
		case t == "pod":
			header = append(header,
				"Pod Capacity", "Pod(%)",
			)
		default:
			header = append(header,
				"CPU USE", "CPU REQ", "CPU REQ(%)", "CPU LIM", "CPU LIM(%)",
				"MEM USE", "MEM REQ", "MEM REQ(%)", "MEM LIM", "MEM LIM(%)",
				"NVIDIA/GPU REQ", "NVIDIA/GPU REQ(%)", "NVIDIA/GPU LIM", "NVIDIA/GPU LIM(%)",
				"PodCount (%)",
			)
		}
	}
	table.SetHeader(header)
	for _, i := range data {
		table.Append(i)
	}
	table.Render()

}

//PodWrite
func PodWrite(data [][]string, resourceType []string, outType bool) {
	table := table(outType)
	var header []string
	header = append(header, "NAMESPACE", "POD NAME", "NODE")

	for _, t := range resourceType {
		switch {
		case t == "cpu":
			header = append(header,
				"CPU USE", "CPU USE(%)", "CPU REQ", "CPU LIM",
			)
		case t == "memory":
			header = append(header,
				"MEM USE", "MEM USE(%)", "MEM REQ", "MEM LIM",
			)
		case t == "gpu":
			header = append(header,
				"GPU REQ", "GPU REQ(%)", "GPU LIM", "GPU LIM(%)", "GPU MODEL",
			)
		default:
			header = append(header,
				"CPU USE ", "CPU USE(%)", "CPU REQ", "CPU LIM",
				"MEM USE", "MEM USE(%)", "MEM REQ", "MEM LIM",
				"GPU REQ", "GPU REQ(%)", "GPU LIM", "GPU LIM(%)", "GPU MODEL",
			)
		}
	}
	table.SetHeader(header)
	for _, i := range data {
		table.Append(i)
	}
	table.Render()

}

//table
func table(outType bool) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	if outType {
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetBorder(false)
		table.SetTablePadding("\t") // pad with tabs
		table.SetNoWhiteSpace(true)
		return table
	}
	return table
}
