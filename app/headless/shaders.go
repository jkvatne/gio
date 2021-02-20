// Code generated by build.go. DO NOT EDIT.

package headless

import "gioui.org/gpu/backend"

var (
	shader_input_vert = backend.ShaderSources{
		Name:      "input.vert",
		Inputs:    []backend.InputLocation{{Name: "position", Location: 0, Semantic: "POSITION", SemanticIndex: 0, Type: 0x0, Size: 4}},
		GLSL100ES: "\nattribute vec4 position;\n\nvoid main()\n{\n    gl_Position = position;\n}\n\n",
		GLSL300ES: "#version 300 es\n\nlayout(location = 0) in vec4 position;\n\nvoid main()\n{\n    gl_Position = position;\n}\n\n",
		GLSL130:   "#version 130\n\nin vec4 position;\n\nvoid main()\n{\n    gl_Position = position;\n}\n\n",
		GLSL150:   "#version 150\n\nin vec4 position;\n\nvoid main()\n{\n    gl_Position = position;\n}\n\n",
		HLSL:      []byte{0x44, 0x58, 0x42, 0x43, 0x35, 0xe9, 0xae, 0x29, 0x96, 0x7e, 0x7c, 0xe6, 0x40, 0xb0, 0x4e, 0x29, 0xd9, 0x98, 0x51, 0x7c, 0x1, 0x0, 0x0, 0x0, 0x10, 0x2, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x38, 0x0, 0x0, 0x0, 0x9c, 0x0, 0x0, 0x0, 0xe0, 0x0, 0x0, 0x0, 0x5c, 0x1, 0x0, 0x0, 0xa8, 0x1, 0x0, 0x0, 0xdc, 0x1, 0x0, 0x0, 0x41, 0x6f, 0x6e, 0x39, 0x5c, 0x0, 0x0, 0x0, 0x5c, 0x0, 0x0, 0x0, 0x0, 0x2, 0xfe, 0xff, 0x34, 0x0, 0x0, 0x0, 0x28, 0x0, 0x0, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x0, 0x24, 0x0, 0x1, 0x0, 0x24, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0xfe, 0xff, 0x1f, 0x0, 0x0, 0x2, 0x5, 0x0, 0x0, 0x80, 0x0, 0x0, 0xf, 0x90, 0x4, 0x0, 0x0, 0x4, 0x0, 0x0, 0x3, 0xc0, 0x0, 0x0, 0xff, 0x90, 0x0, 0x0, 0xe4, 0xa0, 0x0, 0x0, 0xe4, 0x90, 0x1, 0x0, 0x0, 0x2, 0x0, 0x0, 0xc, 0xc0, 0x0, 0x0, 0xe4, 0x90, 0xff, 0xff, 0x0, 0x0, 0x53, 0x48, 0x44, 0x52, 0x3c, 0x0, 0x0, 0x0, 0x40, 0x0, 0x1, 0x0, 0xf, 0x0, 0x0, 0x0, 0x5f, 0x0, 0x0, 0x3, 0xf2, 0x10, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x67, 0x0, 0x0, 0x4, 0xf2, 0x20, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x36, 0x0, 0x0, 0x5, 0xf2, 0x20, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x46, 0x1e, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3e, 0x0, 0x0, 0x1, 0x53, 0x54, 0x41, 0x54, 0x74, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0x44, 0x45, 0x46, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x0, 0x0, 0x4, 0xfe, 0xff, 0x0, 0x1, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x0, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x28, 0x52, 0x29, 0x20, 0x48, 0x4c, 0x53, 0x4c, 0x20, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x20, 0x31, 0x30, 0x2e, 0x31, 0x0, 0x49, 0x53, 0x47, 0x4e, 0x2c, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0xf, 0x0, 0x0, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x0, 0xab, 0xab, 0xab, 0x4f, 0x53, 0x47, 0x4e, 0x2c, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x0, 0x0, 0x0, 0x53, 0x56, 0x5f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x0},
	}
	shader_simple_frag = backend.ShaderSources{
		Name:      "simple.frag",
		GLSL100ES: "precision mediump float;\nprecision highp int;\n\nvoid main()\n{\n    gl_FragData[0] = vec4(0.25, 0.550000011920928955078125, 0.75, 1.0);\n}\n\n",
		GLSL300ES: "#version 300 es\nprecision mediump float;\nprecision highp int;\n\nlayout(location = 0) out vec4 fragColor;\n\nvoid main()\n{\n    fragColor = vec4(0.25, 0.550000011920928955078125, 0.75, 1.0);\n}\n\n",
		GLSL130:   "#version 130\n\nout vec4 fragColor;\n\nvoid main()\n{\n    fragColor = vec4(0.25, 0.550000011920928955078125, 0.75, 1.0);\n}\n\n",
		GLSL150:   "#version 150\n\nout vec4 fragColor;\n\nvoid main()\n{\n    fragColor = vec4(0.25, 0.550000011920928955078125, 0.75, 1.0);\n}\n\n",
		HLSL:      []byte{0x44, 0x58, 0x42, 0x43, 0xf5, 0x46, 0xde, 0x66, 0x24, 0x29, 0xa8, 0xbb, 0x56, 0xea, 0x73, 0xb5, 0x6b, 0x73, 0x12, 0x72, 0x1, 0x0, 0x0, 0x0, 0xdc, 0x1, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x38, 0x0, 0x0, 0x0, 0x90, 0x0, 0x0, 0x0, 0xd0, 0x0, 0x0, 0x0, 0x4c, 0x1, 0x0, 0x0, 0x98, 0x1, 0x0, 0x0, 0xa8, 0x1, 0x0, 0x0, 0x41, 0x6f, 0x6e, 0x39, 0x50, 0x0, 0x0, 0x0, 0x50, 0x0, 0x0, 0x0, 0x0, 0x2, 0xff, 0xff, 0x2c, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x0, 0x24, 0x0, 0x0, 0x2, 0xff, 0xff, 0x51, 0x0, 0x0, 0x5, 0x0, 0x0, 0xf, 0xa0, 0x0, 0x0, 0x80, 0x3e, 0xcd, 0xcc, 0xc, 0x3f, 0x0, 0x0, 0x40, 0x3f, 0x0, 0x0, 0x80, 0x3f, 0x1, 0x0, 0x0, 0x2, 0x0, 0x8, 0xf, 0x80, 0x0, 0x0, 0xe4, 0xa0, 0xff, 0xff, 0x0, 0x0, 0x53, 0x48, 0x44, 0x52, 0x38, 0x0, 0x0, 0x0, 0x40, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0x65, 0x0, 0x0, 0x3, 0xf2, 0x20, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x36, 0x0, 0x0, 0x8, 0xf2, 0x20, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x40, 0x0, 0x0, 0x0, 0x0, 0x80, 0x3e, 0xcd, 0xcc, 0xc, 0x3f, 0x0, 0x0, 0x40, 0x3f, 0x0, 0x0, 0x80, 0x3f, 0x3e, 0x0, 0x0, 0x1, 0x53, 0x54, 0x41, 0x54, 0x74, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0x44, 0x45, 0x46, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x0, 0x0, 0x4, 0xff, 0xff, 0x0, 0x1, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x0, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x28, 0x52, 0x29, 0x20, 0x48, 0x4c, 0x53, 0x4c, 0x20, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x20, 0x31, 0x30, 0x2e, 0x31, 0x0, 0x49, 0x53, 0x47, 0x4e, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x4f, 0x53, 0x47, 0x4e, 0x2c, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x0, 0x0, 0x0, 0x53, 0x56, 0x5f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x0, 0xab, 0xab},
	}
	shader_simple_vert = backend.ShaderSources{
		Name:      "simple.vert",
		GLSL100ES: "\nvoid main()\n{\n    float x;\n    float y;\n    if (gl_VertexID == 0)\n    {\n        x = 0.0;\n        y = 0.5;\n    }\n    else\n    {\n        if (gl_VertexID == 1)\n        {\n            x = 0.5;\n            y = -0.5;\n        }\n        else\n        {\n            x = -0.5;\n            y = -0.5;\n        }\n    }\n    gl_Position = vec4(x, y, 0.5, 1.0);\n}\n\n",
		GLSL300ES: "#version 300 es\n\nvoid main()\n{\n    float x;\n    float y;\n    if (gl_VertexID == 0)\n    {\n        x = 0.0;\n        y = 0.5;\n    }\n    else\n    {\n        if (gl_VertexID == 1)\n        {\n            x = 0.5;\n            y = -0.5;\n        }\n        else\n        {\n            x = -0.5;\n            y = -0.5;\n        }\n    }\n    gl_Position = vec4(x, y, 0.5, 1.0);\n}\n\n",
		GLSL130:   "#version 130\n\nvoid main()\n{\n    float x;\n    float y;\n    if (gl_VertexID == 0)\n    {\n        x = 0.0;\n        y = 0.5;\n    }\n    else\n    {\n        if (gl_VertexID == 1)\n        {\n            x = 0.5;\n            y = -0.5;\n        }\n        else\n        {\n            x = -0.5;\n            y = -0.5;\n        }\n    }\n    gl_Position = vec4(x, y, 0.5, 1.0);\n}\n\n",
		GLSL150:   "#version 150\n\nvoid main()\n{\n    float x;\n    float y;\n    if (gl_VertexID == 0)\n    {\n        x = 0.0;\n        y = 0.5;\n    }\n    else\n    {\n        if (gl_VertexID == 1)\n        {\n            x = 0.5;\n            y = -0.5;\n        }\n        else\n        {\n            x = -0.5;\n            y = -0.5;\n        }\n    }\n    gl_Position = vec4(x, y, 0.5, 1.0);\n}\n\n",
		HLSL:      []byte{0x44, 0x58, 0x42, 0x43, 0xc8, 0x20, 0x5c, 0x22, 0xec, 0xe9, 0xb2, 0x29, 0x40, 0xdf, 0x7c, 0x5a, 0x28, 0xea, 0xc, 0xb8, 0x1, 0x0, 0x0, 0x0, 0x48, 0x2, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x34, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0xb4, 0x0, 0x0, 0x0, 0xe8, 0x0, 0x0, 0x0, 0xcc, 0x1, 0x0, 0x0, 0x52, 0x44, 0x45, 0x46, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x0, 0x0, 0x4, 0xfe, 0xff, 0x0, 0x1, 0x0, 0x0, 0x1c, 0x0, 0x0, 0x0, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x20, 0x28, 0x52, 0x29, 0x20, 0x48, 0x4c, 0x53, 0x4c, 0x20, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x20, 0x31, 0x30, 0x2e, 0x31, 0x0, 0x49, 0x53, 0x47, 0x4e, 0x2c, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x0, 0x0, 0x53, 0x56, 0x5f, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x49, 0x44, 0x0, 0x4f, 0x53, 0x47, 0x4e, 0x2c, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x0, 0x0, 0x0, 0x53, 0x56, 0x5f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x0, 0x53, 0x48, 0x44, 0x52, 0xdc, 0x0, 0x0, 0x0, 0x40, 0x0, 0x1, 0x0, 0x37, 0x0, 0x0, 0x0, 0x60, 0x0, 0x0, 0x4, 0x12, 0x10, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x67, 0x0, 0x0, 0x4, 0xf2, 0x20, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x68, 0x0, 0x0, 0x2, 0x1, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x7, 0x12, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x10, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x40, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x37, 0x0, 0x0, 0xf, 0x32, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0x0, 0x0, 0xbf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0xbf, 0x0, 0x0, 0x0, 0xbf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x37, 0x0, 0x0, 0xc, 0x32, 0x20, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x10, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x46, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x36, 0x0, 0x0, 0x8, 0xc2, 0x20, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3f, 0x0, 0x0, 0x80, 0x3f, 0x3e, 0x0, 0x0, 0x1, 0x53, 0x54, 0x41, 0x54, 0x74, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	}
)
