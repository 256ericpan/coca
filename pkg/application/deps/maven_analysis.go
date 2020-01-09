package deps

import (
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/infrastructure/xmlparse"
	"os"
)

func AnalysisMaven(xmlPath string) []domain.JDependency {
	xmlFile, _ := os.Open(xmlPath)
	parseXml := xmlparse.ParseXml(xmlFile)
	for _, element := range parseXml.Elements {
		val := element.Val.(xmlparse.XmlNode)
		if val.Name == "dependencies" {
			return BuildDeps(val)
		}
	}
	return nil
}

func BuildDeps(val xmlparse.XmlNode) []domain.JDependency {
	var deps []domain.JDependency = nil
	for _, depElement := range val.Elements {
		depNode := depElement.Val.(xmlparse.XmlNode)
		dependency := domain.NewJDependency("", "")

		for _, depValue := range depNode.Elements {
			node := depValue.Val.(xmlparse.XmlNode)
			if node.Name == "groupId" {
				for _, textNode := range node.Elements {
					dependency.GroupId = textNode.Val.(string)
				}
			}

			if node.Name == "artifactId" {
				for _, textNode := range node.Elements {
					dependency.ArtifactId = textNode.Val.(string)
				}
			}

			if node.Name == "scope" {
				for _, textNode := range node.Elements {
					dependency.Scope = textNode.Val.(string)
				}
			}
		}

		deps = append(deps, *dependency)
	}

	return deps
}

