package template

func TemplateGridTemplateAreasHTML() string {
	gridTemplateAreasTemplate := `
          grid-template-areas:
            "tn  .  .  ."
            ".   p1 p2 p3"
            "tv1 a  b  c"
            "tv2 d  e  f"
            "tv3 g  h  i";
	`
	return gridTemplateAreasTemplate
}

func TemplateVariantXPod() string {
	variantXPodTemplate := `
          <div class="square {{index .Assigns.Data 0 0}}" style="grid-area: a;">
          </div>
	`
	return variantXPodTemplate
}

func TemplatePod() string {
	PodTemplate := `
          <div class="vertical-txt" style="grid-area: p1;">
            <p>pod1</p>
          </div>
	`
	return PodTemplate
}

func TemplateTestVariant() string {
	testVariantTemplate := `
          <div class="square-txt" style="grid-area: tv1;">
            <p>test variant 1</p>
          </div>
	`
	return testVariantTemplate
}

func TemplateTest(testName string) string {
	testTemplate := `
        <div class="{{ $.testName }}">
          <div style="grid-area: tn;">
            <p>test 1</p>
          </div>
          {{- range $v := TemplatePod }}
            {{ $v }}
          {{- end }}
          {{- range $v := TemplateTestVariant }}
            {{ $v }}
          {{- end }}
          {{- range $v := TemplateVariantXPod }}
            {{ $v }}
          {{- end }}
        </div>
	`
	return testTemplate
}
