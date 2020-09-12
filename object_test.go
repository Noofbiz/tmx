package tmx

import (
	"os"
	"testing"
)

func TestObjectExternal(t *testing.T) {
	TMXURL = "testData/objects.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse object data")
		return
	}
	exp := struct {
		objGroup ObjectGroup
		group    Group
	}{
		objGroup: ObjectGroup{
			Objects: []Object{
				Object{
					ID:   1,
					Name: "Rectangle",
				},
				Object{
					ID:     3,
					X:      26,
					Y:      5,
					Width:  15,
					Height: 40,
				},
				Object{
					ID:     7,
					X:      0,
					Y:      0,
					Width:  15,
					Height: 40,
				},
			},
			Name: "Object Layer 1",
		},
		group: Group{
			ImageLayers: []ImageLayer{
				ImageLayer{
					Name:    "Image Layer 1",
					OffsetX: 5,
					OffsetY: 5,
				},
			},
			Name:    "Group 1",
			OffsetX: 2,
			OffsetY: 2,
		},
	}
	for _, og := range m.ObjectGroups {
		if og.Name != exp.objGroup.Name {
			t.Errorf("Unexpected object group name\nWanted: %v\nGot: %v", exp.objGroup.Name, og.Name)
			return
		}
		for i, obj := range og.Objects {
			if obj.ID == 1 {
				if obj.Name != exp.objGroup.Objects[i].Name {
					t.Error("Object 1 was not named Rectangle")
					return
				}
			} else if obj.ID == 3 || obj.ID == 7 {
				if len(obj.Ellipses) != 1 {
					t.Errorf("Object %v did not contain an ellipse", obj.ID)
					return
				}
				if obj.Width != exp.objGroup.Objects[i].Width {
					t.Errorf("Object %v did not properly get width from the template\nWanted: %v\nGot: %v", obj.ID, exp.objGroup.Objects[i].Width, obj.Width)
					return
				}
				if obj.Height != exp.objGroup.Objects[i].Height {
					t.Errorf("Object %v did not properly get height from the template\nWanted: %v\nGot: %v", obj.ID, exp.objGroup.Objects[i].Height, obj.Height)
					return
				}
				if obj.X != exp.objGroup.Objects[i].X {
					t.Errorf("Object %v did not properly get x from the template\nWanted: %v\nGot: %v", obj.ID, exp.objGroup.Objects[i].X, obj.X)
					return
				}
				if obj.Y != exp.objGroup.Objects[i].Y {
					t.Errorf("Object %v did not properly get y from the template\nWanted: %v\nGot: %v", obj.ID, exp.objGroup.Objects[i].Y, obj.Y)
					return
				}
			} else {
				t.Errorf("Unexpected object ID: %v", obj.ID)
				return
			}
		}
	}
	if m.Groups[0].Name != exp.group.Name {
		t.Errorf("Group name was not properly set\nWanted: %v\nGot: %v", exp.group.Name, m.Groups[0].Name)
		return
	}
	if m.Groups[0].OffsetX != exp.group.OffsetX {
		t.Errorf("Group offset X was not properly set\nWanted: %v\nGot: %v", exp.group.OffsetX, m.Groups[0].OffsetX)
		return
	}
	if m.Groups[0].OffsetY != exp.group.OffsetY {
		t.Errorf("Group offset Y was not properly set\nWanted: %v\nGot: %v", exp.group.OffsetY, m.Groups[0].OffsetY)
		return
	}
	if m.Groups[0].ImageLayers[0].Name != exp.group.ImageLayers[0].Name {
		t.Errorf("Group's image layer was not properly set\nWanted: %v\nGot: %v", exp.group.ImageLayers[0].Name, m.Groups[0].ImageLayers[0].Name)
		return
	}
}

func TestObjectTemplateNotExist(t *testing.T) {
	TMXURL = "testData/objectTemplateNotExist.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Error("Able to parse with non existant template file")
		return
	}
}

func TestObjectMalformed(t *testing.T) {
	TMXURL = "testData/malformedObject.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Error("Able to parse with malformed object elements")
		return
	}
}

func TestObjectTemplateMalformed(t *testing.T) {
	TMXURL = "testData/malformedObjectTemplate.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Error("Able to parse with malformed object template")
		return
	}
}

func TestImageLayerMalformed(t *testing.T) {
	TMXURL = "testData/malformedImageLayer.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Error("Able to parse with malformed image layer")
		return
	}
}

func TestText(t *testing.T) {
	TMXURL = "testData/text.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unble to parse text. Error was %v", err)
		return
	}
	if m.ObjectGroups[0].Objects[0].Text[0].CharData != "Hello World" {
		t.Errorf("Did not parse text correctly\nWanted: %v\nGot: %v", "Hello World", m.ObjectGroups[0].Objects[0].Text[0].CharData)
		return
	}
}

func TestTextMalformed(t *testing.T) {
	TMXURL = "testData/malformedText.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Error("Able to parse with malformed text")
		return
	}
}

func TestPropertyParsed(t *testing.T) {
	TMXURL = "testData/properties.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Error("Unable to parse object data")
	}
	prop1 := m.ObjectGroups[0].Objects[0].Properties[0]
	if prop1.Name != "attrValue" {
		t.Error("Unable to parse object name")
	}
	if prop1.Value != "This is an attribute value" {
		t.Error("Unable to parse object value")
	}

	prop2 := m.ObjectGroups[0].Objects[0].Properties[1]
	if prop2.Name != "multilineValue" {
		t.Error("Unable to parse object name")
	}
	if prop2.Value != `This is
a multiline value` {
		t.Error("Unable to parse object value")
	}

}
