// Code generated by build.go. DO NOT EDIT.

package gpu

import "gioui.org/gpu/backend"

var (
	shader_blit_frag = [...]backend.ShaderSources{
		backend.ShaderSources{
			Uniforms: backend.UniformsReflection{
				Blocks:    []backend.UniformBlock{backend.UniformBlock{Name: "Color", Binding: 0}},
				Locations: []backend.UniformLocation{backend.UniformLocation{Name: "_12._color", Type: 0x0, Size: 4, Offset: 0}},
				Size:      16,
			},
			GLSL100ES: "#version 100\nprecision mediump float;\nprecision highp int;\n\nstruct Color\n{\n    vec4 _color;\n};\n\nuniform Color _12;\n\nvarying vec2 vUV;\n\nvoid main()\n{\n    gl_FragData[0] = _12._color;\n}\n\n",
			GLSL300ES: "#version 300 es\nprecision mediump float;\nprecision highp int;\n\nlayout(std140) uniform Color\n{\n    vec4 _color;\n} _12;\n\nlayout(location = 0) out vec4 fragColor;\nin vec2 vUV;\n\nvoid main()\n{\n    fragColor = _12._color;\n}\n\n",
			/*
			   cbuffer Color : register(b0)
			   {
			       float4 _12_color : packoffset(c0);
			   };


			   static float4 fragColor;
			   static float2 vUV;

			   struct SPIRV_Cross_Input
			   {
			       float2 vUV : TEXCOORD0;
			   };

			   struct SPIRV_Cross_Output
			   {
			       float4 fragColor : SV_Target0;
			   };

			   void frag_main()
			   {
			       fragColor = _12_color;
			   }

			   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
			   {
			       vUV = stage_input.vUV;
			       frag_main();
			       SPIRV_Cross_Output stage_output;
			       stage_output.fragColor = fragColor;
			       return stage_output;
			   }

			*/
			HLSL: []byte(nil),
		},
		backend.ShaderSources{
			Textures:  []backend.TextureBinding{backend.TextureBinding{Name: "tex", Binding: 0}},
			GLSL100ES: "#version 100\nprecision mediump float;\nprecision highp int;\n\nuniform mediump sampler2D tex;\n\nvarying vec2 vUV;\n\nvoid main()\n{\n    gl_FragData[0] = texture2D(tex, vUV);\n}\n\n",
			GLSL300ES: "#version 300 es\nprecision mediump float;\nprecision highp int;\n\nuniform mediump sampler2D tex;\n\nlayout(location = 0) out vec4 fragColor;\nin vec2 vUV;\n\nvoid main()\n{\n    fragColor = texture(tex, vUV);\n}\n\n",
			/*
			   Texture2D<float4> tex : register(t0);
			   SamplerState _tex_sampler : register(s0);

			   static float4 fragColor;
			   static float2 vUV;

			   struct SPIRV_Cross_Input
			   {
			       float2 vUV : TEXCOORD0;
			   };

			   struct SPIRV_Cross_Output
			   {
			       float4 fragColor : SV_Target0;
			   };

			   void frag_main()
			   {
			       fragColor = tex.Sample(_tex_sampler, vUV);
			   }

			   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
			   {
			       vUV = stage_input.vUV;
			       frag_main();
			       SPIRV_Cross_Output stage_output;
			       stage_output.fragColor = fragColor;
			       return stage_output;
			   }

			*/
			HLSL: []byte(nil),
		},
	}
	shader_blit_vert = backend.ShaderSources{
		Inputs: []backend.InputLocation{backend.InputLocation{Name: "pos", Location: 0, Semantic: "POSITION", SemanticIndex: 0, Type: 0x0, Size: 2}, backend.InputLocation{Name: "uv", Location: 1, Semantic: "NORMAL", SemanticIndex: 0, Type: 0x0, Size: 2}},
		Uniforms: backend.UniformsReflection{
			Blocks:    []backend.UniformBlock{backend.UniformBlock{Name: "Block", Binding: 0}},
			Locations: []backend.UniformLocation{backend.UniformLocation{Name: "_24.transform", Type: 0x0, Size: 4, Offset: 0}, backend.UniformLocation{Name: "_24.uvTransform", Type: 0x0, Size: 4, Offset: 16}, backend.UniformLocation{Name: "_24.z", Type: 0x0, Size: 1, Offset: 32}},
			Size:      36,
		},
		GLSL100ES: "#version 100\n\nstruct Block\n{\n    vec4 transform;\n    vec4 uvTransform;\n    float z;\n};\n\nuniform Block _24;\n\nattribute vec2 pos;\nvarying vec2 vUV;\nattribute vec2 uv;\n\nvec4 toClipSpace(vec4 pos_1)\n{\n    return pos_1;\n}\n\nvoid main()\n{\n    vec2 p = (pos * _24.transform.xy) + _24.transform.zw;\n    vec4 param = vec4(p, _24.z, 1.0);\n    gl_Position = toClipSpace(param);\n    vUV = (uv * _24.uvTransform.xy) + _24.uvTransform.zw;\n}\n\n",
		GLSL300ES: "#version 300 es\n\nlayout(std140) uniform Block\n{\n    vec4 transform;\n    vec4 uvTransform;\n    float z;\n} _24;\n\nlayout(location = 0) in vec2 pos;\nout vec2 vUV;\nlayout(location = 1) in vec2 uv;\n\nvec4 toClipSpace(vec4 pos_1)\n{\n    return pos_1;\n}\n\nvoid main()\n{\n    vec2 p = (pos * _24.transform.xy) + _24.transform.zw;\n    vec4 param = vec4(p, _24.z, 1.0);\n    gl_Position = toClipSpace(param);\n    vUV = (uv * _24.uvTransform.xy) + _24.uvTransform.zw;\n}\n\n",
		/*
		   cbuffer Block : register(b0)
		   {
		       float4 _24_transform : packoffset(c0);
		       float4 _24_uvTransform : packoffset(c1);
		       float _24_z : packoffset(c2);
		   };


		   static float4 gl_Position;
		   static float2 pos;
		   static float2 vUV;
		   static float2 uv;

		   struct SPIRV_Cross_Input
		   {
		       float2 pos : POSITION;
		       float2 uv : NORMAL;
		   };

		   struct SPIRV_Cross_Output
		   {
		       float2 vUV : TEXCOORD0;
		       float4 gl_Position : SV_Position;
		   };

		   float4 toClipSpace(float4 pos_1)
		   {
		       return pos_1;
		   }

		   void vert_main()
		   {
		       float2 p = (pos * _24_transform.xy) + _24_transform.zw;
		       float4 param = float4(p, _24_z, 1.0f);
		       gl_Position = toClipSpace(param);
		       vUV = (uv * _24_uvTransform.xy) + _24_uvTransform.zw;
		   }

		   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
		   {
		       pos = stage_input.pos;
		       uv = stage_input.uv;
		       vert_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.gl_Position = gl_Position;
		       stage_output.vUV = vUV;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
	shader_cover_frag = [...]backend.ShaderSources{
		backend.ShaderSources{
			Uniforms: backend.UniformsReflection{
				Blocks:    []backend.UniformBlock{backend.UniformBlock{Name: "Color", Binding: 0}},
				Locations: []backend.UniformLocation{backend.UniformLocation{Name: "_12._color", Type: 0x0, Size: 4, Offset: 0}},
				Size:      16,
			},
			Textures:  []backend.TextureBinding{backend.TextureBinding{Name: "cover", Binding: 1}},
			GLSL100ES: "#version 100\nprecision mediump float;\nprecision highp int;\n\nstruct Color\n{\n    vec4 _color;\n};\n\nuniform Color _12;\n\nuniform mediump sampler2D cover;\n\nvarying highp vec2 vCoverUV;\nvarying vec2 vUV;\n\nvoid main()\n{\n    gl_FragData[0] = _12._color;\n    float cover_1 = abs(texture2D(cover, vCoverUV).x);\n    gl_FragData[0] *= cover_1;\n}\n\n",
			GLSL300ES: "#version 300 es\nprecision mediump float;\nprecision highp int;\n\nlayout(std140) uniform Color\n{\n    vec4 _color;\n} _12;\n\nuniform mediump sampler2D cover;\n\nlayout(location = 0) out vec4 fragColor;\nin highp vec2 vCoverUV;\nin vec2 vUV;\n\nvoid main()\n{\n    fragColor = _12._color;\n    float cover_1 = abs(texture(cover, vCoverUV).x);\n    fragColor *= cover_1;\n}\n\n",
			/*
			   cbuffer Color : register(b0)
			   {
			       float4 _12_color : packoffset(c0);
			   };

			   Texture2D<float4> cover : register(t1);
			   SamplerState _cover_sampler : register(s1);

			   static float4 fragColor;
			   static float2 vCoverUV;
			   static float2 vUV;

			   struct SPIRV_Cross_Input
			   {
			       float2 vCoverUV : TEXCOORD0;
			       float2 vUV : TEXCOORD1;
			   };

			   struct SPIRV_Cross_Output
			   {
			       float4 fragColor : SV_Target0;
			   };

			   void frag_main()
			   {
			       fragColor = _12_color;
			       float cover_1 = abs(cover.Sample(_cover_sampler, vCoverUV).x);
			       fragColor *= cover_1;
			   }

			   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
			   {
			       vCoverUV = stage_input.vCoverUV;
			       vUV = stage_input.vUV;
			       frag_main();
			       SPIRV_Cross_Output stage_output;
			       stage_output.fragColor = fragColor;
			       return stage_output;
			   }

			*/
			HLSL: []byte(nil),
		},
		backend.ShaderSources{
			Textures:  []backend.TextureBinding{backend.TextureBinding{Name: "tex", Binding: 0}, backend.TextureBinding{Name: "cover", Binding: 1}},
			GLSL100ES: "#version 100\nprecision mediump float;\nprecision highp int;\n\nuniform mediump sampler2D tex;\nuniform mediump sampler2D cover;\n\nvarying vec2 vUV;\nvarying highp vec2 vCoverUV;\n\nvoid main()\n{\n    gl_FragData[0] = texture2D(tex, vUV);\n    float cover_1 = abs(texture2D(cover, vCoverUV).x);\n    gl_FragData[0] *= cover_1;\n}\n\n",
			GLSL300ES: "#version 300 es\nprecision mediump float;\nprecision highp int;\n\nuniform mediump sampler2D tex;\nuniform mediump sampler2D cover;\n\nlayout(location = 0) out vec4 fragColor;\nin vec2 vUV;\nin highp vec2 vCoverUV;\n\nvoid main()\n{\n    fragColor = texture(tex, vUV);\n    float cover_1 = abs(texture(cover, vCoverUV).x);\n    fragColor *= cover_1;\n}\n\n",
			/*
			   Texture2D<float4> tex : register(t0);
			   SamplerState _tex_sampler : register(s0);
			   Texture2D<float4> cover : register(t1);
			   SamplerState _cover_sampler : register(s1);

			   static float4 fragColor;
			   static float2 vUV;
			   static float2 vCoverUV;

			   struct SPIRV_Cross_Input
			   {
			       float2 vCoverUV : TEXCOORD0;
			       float2 vUV : TEXCOORD1;
			   };

			   struct SPIRV_Cross_Output
			   {
			       float4 fragColor : SV_Target0;
			   };

			   void frag_main()
			   {
			       fragColor = tex.Sample(_tex_sampler, vUV);
			       float cover_1 = abs(cover.Sample(_cover_sampler, vCoverUV).x);
			       fragColor *= cover_1;
			   }

			   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
			   {
			       vUV = stage_input.vUV;
			       vCoverUV = stage_input.vCoverUV;
			       frag_main();
			       SPIRV_Cross_Output stage_output;
			       stage_output.fragColor = fragColor;
			       return stage_output;
			   }

			*/
			HLSL: []byte(nil),
		},
	}
	shader_cover_vert = backend.ShaderSources{
		Inputs: []backend.InputLocation{backend.InputLocation{Name: "pos", Location: 0, Semantic: "POSITION", SemanticIndex: 0, Type: 0x0, Size: 2}, backend.InputLocation{Name: "uv", Location: 1, Semantic: "NORMAL", SemanticIndex: 0, Type: 0x0, Size: 2}},
		Uniforms: backend.UniformsReflection{
			Blocks:    []backend.UniformBlock{backend.UniformBlock{Name: "Block", Binding: 0}},
			Locations: []backend.UniformLocation{backend.UniformLocation{Name: "_66.transform", Type: 0x0, Size: 4, Offset: 0}, backend.UniformLocation{Name: "_66.uvCoverTransform", Type: 0x0, Size: 4, Offset: 16}, backend.UniformLocation{Name: "_66.uvTransform", Type: 0x0, Size: 4, Offset: 32}, backend.UniformLocation{Name: "_66.z", Type: 0x0, Size: 1, Offset: 48}},
			Size:      52,
		},
		GLSL100ES: "#version 100\n\nstruct Block\n{\n    vec4 transform;\n    vec4 uvCoverTransform;\n    vec4 uvTransform;\n    float z;\n};\n\nuniform Block _66;\n\nattribute vec2 pos;\nvarying vec2 vUV;\nattribute vec2 uv;\nvarying vec2 vCoverUV;\n\nvec4 toClipSpace(vec4 pos_1)\n{\n    return pos_1;\n}\n\nvec3[2] fboTextureTransform()\n{\n    vec3 t[2];\n    t[0] = vec3(1.0, 0.0, 0.0);\n    t[1] = vec3(0.0, 1.0, 0.0);\n    return t;\n}\n\nvec3 transform3x2(vec3 t[2], vec3 v)\n{\n    return vec3(dot(t[0], v), dot(t[1], v), dot(vec3(0.0, 0.0, 1.0), v));\n}\n\nvoid main()\n{\n    vec4 param = vec4((pos * _66.transform.xy) + _66.transform.zw, _66.z, 1.0);\n    gl_Position = toClipSpace(param);\n    vUV = (uv * _66.uvTransform.xy) + _66.uvTransform.zw;\n    vec3 fboTrans[2] = fboTextureTransform();\n    vec3 param_1[2] = fboTrans;\n    vec3 param_2 = vec3(uv, 1.0);\n    vec3 uv3 = transform3x2(param_1, param_2);\n    vCoverUV = ((uv3 * vec3(_66.uvCoverTransform.xy, 1.0)) + vec3(_66.uvCoverTransform.zw, 0.0)).xy;\n}\n\n",
		GLSL300ES: "#version 300 es\n\nlayout(std140) uniform Block\n{\n    vec4 transform;\n    vec4 uvCoverTransform;\n    vec4 uvTransform;\n    float z;\n} _66;\n\nlayout(location = 0) in vec2 pos;\nout vec2 vUV;\nlayout(location = 1) in vec2 uv;\nout vec2 vCoverUV;\n\nvec4 toClipSpace(vec4 pos_1)\n{\n    return pos_1;\n}\n\nvec3[2] fboTextureTransform()\n{\n    vec3 t[2];\n    t[0] = vec3(1.0, 0.0, 0.0);\n    t[1] = vec3(0.0, 1.0, 0.0);\n    return t;\n}\n\nvec3 transform3x2(vec3 t[2], vec3 v)\n{\n    return vec3(dot(t[0], v), dot(t[1], v), dot(vec3(0.0, 0.0, 1.0), v));\n}\n\nvoid main()\n{\n    vec4 param = vec4((pos * _66.transform.xy) + _66.transform.zw, _66.z, 1.0);\n    gl_Position = toClipSpace(param);\n    vUV = (uv * _66.uvTransform.xy) + _66.uvTransform.zw;\n    vec3 fboTrans[2] = fboTextureTransform();\n    vec3 param_1[2] = fboTrans;\n    vec3 param_2 = vec3(uv, 1.0);\n    vec3 uv3 = transform3x2(param_1, param_2);\n    vCoverUV = ((uv3 * vec3(_66.uvCoverTransform.xy, 1.0)) + vec3(_66.uvCoverTransform.zw, 0.0)).xy;\n}\n\n",
		/*
		   cbuffer Block : register(b0)
		   {
		       float4 _66_transform : packoffset(c0);
		       float4 _66_uvCoverTransform : packoffset(c1);
		       float4 _66_uvTransform : packoffset(c2);
		       float _66_z : packoffset(c3);
		   };


		   static float4 gl_Position;
		   static float2 pos;
		   static float2 vUV;
		   static float2 uv;
		   static float2 vCoverUV;

		   struct SPIRV_Cross_Input
		   {
		       float2 pos : POSITION;
		       float2 uv : NORMAL;
		   };

		   struct SPIRV_Cross_Output
		   {
		       float2 vCoverUV : TEXCOORD0;
		       float2 vUV : TEXCOORD1;
		       float4 gl_Position : SV_Position;
		   };

		   float4 toClipSpace(float4 pos_1)
		   {
		       return pos_1;
		   }

		   void fboTextureTransform(out float3 SPIRV_Cross_return_value[2])
		   {
		       float3 t[2];
		       t[0] = float3(1.0f, 0.0f, 0.0f);
		       t[1] = float3(0.0f, 1.0f, 0.0f);
		       SPIRV_Cross_return_value = t;
		   }

		   float3 transform3x2(float3 t[2], float3 v)
		   {
		       return float3(dot(t[0], v), dot(t[1], v), dot(float3(0.0f, 0.0f, 1.0f), v));
		   }

		   void vert_main()
		   {
		       float4 param = float4((pos * _66_transform.xy) + _66_transform.zw, _66_z, 1.0f);
		       gl_Position = toClipSpace(param);
		       vUV = (uv * _66_uvTransform.xy) + _66_uvTransform.zw;
		       float3 _101[2];
		       fboTextureTransform(_101);
		       float3 fboTrans[2] = _101;
		       float3 param_1[2] = fboTrans;
		       float3 param_2 = float3(uv, 1.0f);
		       float3 uv3 = transform3x2(param_1, param_2);
		       vCoverUV = ((uv3 * float3(_66_uvCoverTransform.xy, 1.0f)) + float3(_66_uvCoverTransform.zw, 0.0f)).xy;
		   }

		   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
		   {
		       pos = stage_input.pos;
		       uv = stage_input.uv;
		       vert_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.gl_Position = gl_Position;
		       stage_output.vUV = vUV;
		       stage_output.vCoverUV = vCoverUV;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
	shader_intersect_frag = backend.ShaderSources{
		Textures:  []backend.TextureBinding{backend.TextureBinding{Name: "cover", Binding: 0}},
		GLSL100ES: "#version 100\nprecision mediump float;\nprecision highp int;\n\nuniform mediump sampler2D cover;\n\nvarying highp vec2 vUV;\n\nvoid main()\n{\n    float cover_1 = abs(texture2D(cover, vUV).x);\n    gl_FragData[0].x = cover_1;\n}\n\n",
		GLSL300ES: "#version 300 es\nprecision mediump float;\nprecision highp int;\n\nuniform mediump sampler2D cover;\n\nin highp vec2 vUV;\nlayout(location = 0) out vec4 fragColor;\n\nvoid main()\n{\n    float cover_1 = abs(texture(cover, vUV).x);\n    fragColor.x = cover_1;\n}\n\n",
		/*
		   Texture2D<float4> cover : register(t0);
		   SamplerState _cover_sampler : register(s0);

		   static float2 vUV;
		   static float4 fragColor;

		   struct SPIRV_Cross_Input
		   {
		       float2 vUV : TEXCOORD0;
		   };

		   struct SPIRV_Cross_Output
		   {
		       float4 fragColor : SV_Target0;
		   };

		   void frag_main()
		   {
		       float cover_1 = abs(cover.Sample(_cover_sampler, vUV).x);
		       fragColor.x = cover_1;
		   }

		   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
		   {
		       vUV = stage_input.vUV;
		       frag_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.fragColor = fragColor;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
	shader_intersect_vert = backend.ShaderSources{
		Inputs: []backend.InputLocation{backend.InputLocation{Name: "pos", Location: 0, Semantic: "POSITION", SemanticIndex: 0, Type: 0x0, Size: 2}, backend.InputLocation{Name: "uv", Location: 1, Semantic: "NORMAL", SemanticIndex: 0, Type: 0x0, Size: 2}},
		Uniforms: backend.UniformsReflection{
			Blocks:    []backend.UniformBlock{backend.UniformBlock{Name: "Block", Binding: 0}},
			Locations: []backend.UniformLocation{backend.UniformLocation{Name: "_40.uvTransform", Type: 0x0, Size: 4, Offset: 0}},
			Size:      16,
		},
		GLSL100ES: "#version 100\n\nstruct Block\n{\n    vec4 uvTransform;\n};\n\nuniform Block _40;\n\nattribute vec2 pos;\nvarying vec2 vUV;\nattribute vec2 uv;\n\nvoid main()\n{\n    vec2 p = pos;\n    p.y = -p.y;\n    gl_Position = vec4(p, 0.0, 1.0);\n    vUV = (uv * _40.uvTransform.xy) + _40.uvTransform.zw;\n}\n\n",
		GLSL300ES: "#version 300 es\n\nlayout(std140) uniform Block\n{\n    vec4 uvTransform;\n} _40;\n\nlayout(location = 0) in vec2 pos;\nout vec2 vUV;\nlayout(location = 1) in vec2 uv;\n\nvoid main()\n{\n    vec2 p = pos;\n    p.y = -p.y;\n    gl_Position = vec4(p, 0.0, 1.0);\n    vUV = (uv * _40.uvTransform.xy) + _40.uvTransform.zw;\n}\n\n",
		/*
		   cbuffer Block : register(b0)
		   {
		       float4 _40_uvTransform : packoffset(c0);
		   };


		   static float4 gl_Position;
		   static float2 pos;
		   static float2 vUV;
		   static float2 uv;

		   struct SPIRV_Cross_Input
		   {
		       float2 pos : POSITION;
		       float2 uv : NORMAL;
		   };

		   struct SPIRV_Cross_Output
		   {
		       float2 vUV : TEXCOORD0;
		       float4 gl_Position : SV_Position;
		   };

		   void vert_main()
		   {
		       float2 p = pos;
		       p.y = -p.y;
		       gl_Position = float4(p, 0.0f, 1.0f);
		       vUV = (uv * _40_uvTransform.xy) + _40_uvTransform.zw;
		   }

		   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
		   {
		       pos = stage_input.pos;
		       uv = stage_input.uv;
		       vert_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.gl_Position = gl_Position;
		       stage_output.vUV = vUV;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
	shader_stencil_frag = backend.ShaderSources{
		GLSL100ES: "#version 100\nprecision mediump float;\nprecision highp int;\n\nvarying vec2 vTo;\nvarying vec2 vFrom;\nvarying vec2 vCtrl;\n\nvoid main()\n{\n    float dx = vTo.x - vFrom.x;\n    bool increasing = vTo.x >= vFrom.x;\n    bvec2 _35 = bvec2(increasing);\n    vec2 left = vec2(_35.x ? vFrom.x : vTo.x, _35.y ? vFrom.y : vTo.y);\n    bvec2 _41 = bvec2(increasing);\n    vec2 right = vec2(_41.x ? vTo.x : vFrom.x, _41.y ? vTo.y : vFrom.y);\n    vec2 extent = clamp(vec2(vFrom.x, vTo.x), vec2(-0.5), vec2(0.5));\n    float midx = mix(extent.x, extent.y, 0.5);\n    float x0 = midx - left.x;\n    vec2 p1 = vCtrl - left;\n    vec2 v = right - vCtrl;\n    float t = x0 / (p1.x + sqrt((p1.x * p1.x) + ((v.x - p1.x) * x0)));\n    float y = mix(mix(left.y, vCtrl.y, t), mix(vCtrl.y, right.y, t), t);\n    vec2 d_half = mix(p1, v, vec2(t));\n    float dy = d_half.y / d_half.x;\n    float width = extent.y - extent.x;\n    dy = abs(dy * width);\n    vec4 sides = vec4((dy * 0.5) + y, (dy * (-0.5)) + y, (0.5 - y) / dy, ((-0.5) - y) / dy);\n    sides = clamp(sides + vec4(0.5), vec4(0.0), vec4(1.0));\n    float area = 0.5 * ((((sides.z - (sides.z * sides.y)) + 1.0) - sides.x) + (sides.x * sides.w));\n    area *= width;\n    if (width == 0.0)\n    {\n        area = 0.0;\n    }\n    gl_FragData[0].x = area;\n}\n\n",
		GLSL300ES: "#version 300 es\nprecision mediump float;\nprecision highp int;\n\nin vec2 vTo;\nin vec2 vFrom;\nin vec2 vCtrl;\nlayout(location = 0) out vec4 fragCover;\n\nvoid main()\n{\n    float dx = vTo.x - vFrom.x;\n    bool increasing = vTo.x >= vFrom.x;\n    bvec2 _35 = bvec2(increasing);\n    vec2 left = vec2(_35.x ? vFrom.x : vTo.x, _35.y ? vFrom.y : vTo.y);\n    bvec2 _41 = bvec2(increasing);\n    vec2 right = vec2(_41.x ? vTo.x : vFrom.x, _41.y ? vTo.y : vFrom.y);\n    vec2 extent = clamp(vec2(vFrom.x, vTo.x), vec2(-0.5), vec2(0.5));\n    float midx = mix(extent.x, extent.y, 0.5);\n    float x0 = midx - left.x;\n    vec2 p1 = vCtrl - left;\n    vec2 v = right - vCtrl;\n    float t = x0 / (p1.x + sqrt((p1.x * p1.x) + ((v.x - p1.x) * x0)));\n    float y = mix(mix(left.y, vCtrl.y, t), mix(vCtrl.y, right.y, t), t);\n    vec2 d_half = mix(p1, v, vec2(t));\n    float dy = d_half.y / d_half.x;\n    float width = extent.y - extent.x;\n    dy = abs(dy * width);\n    vec4 sides = vec4((dy * 0.5) + y, (dy * (-0.5)) + y, (0.5 - y) / dy, ((-0.5) - y) / dy);\n    sides = clamp(sides + vec4(0.5), vec4(0.0), vec4(1.0));\n    float area = 0.5 * ((((sides.z - (sides.z * sides.y)) + 1.0) - sides.x) + (sides.x * sides.w));\n    area *= width;\n    if (width == 0.0)\n    {\n        area = 0.0;\n    }\n    fragCover.x = area;\n}\n\n",
		/*
		   static float2 vTo;
		   static float2 vFrom;
		   static float2 vCtrl;
		   static float4 fragCover;

		   struct SPIRV_Cross_Input
		   {
		       float2 vFrom : TEXCOORD0;
		       float2 vCtrl : TEXCOORD1;
		       float2 vTo : TEXCOORD2;
		   };

		   struct SPIRV_Cross_Output
		   {
		       float4 fragCover : SV_Target0;
		   };

		   void frag_main()
		   {
		       float dx = vTo.x - vFrom.x;
		       bool increasing = vTo.x >= vFrom.x;
		       bool2 _35 = increasing.xx;
		       float2 left = float2(_35.x ? vFrom.x : vTo.x, _35.y ? vFrom.y : vTo.y);
		       bool2 _41 = increasing.xx;
		       float2 right = float2(_41.x ? vTo.x : vFrom.x, _41.y ? vTo.y : vFrom.y);
		       float2 extent = clamp(float2(vFrom.x, vTo.x), (-0.5f).xx, 0.5f.xx);
		       float midx = lerp(extent.x, extent.y, 0.5f);
		       float x0 = midx - left.x;
		       float2 p1 = vCtrl - left;
		       float2 v = right - vCtrl;
		       float t = x0 / (p1.x + sqrt((p1.x * p1.x) + ((v.x - p1.x) * x0)));
		       float y = lerp(lerp(left.y, vCtrl.y, t), lerp(vCtrl.y, right.y, t), t);
		       float2 d_half = lerp(p1, v, t.xx);
		       float dy = d_half.y / d_half.x;
		       float width = extent.y - extent.x;
		       dy = abs(dy * width);
		       float4 sides = float4((dy * 0.5f) + y, (dy * (-0.5f)) + y, (0.5f - y) / dy, ((-0.5f) - y) / dy);
		       sides = clamp(sides + 0.5f.xxxx, 0.0f.xxxx, 1.0f.xxxx);
		       float area = 0.5f * ((((sides.z - (sides.z * sides.y)) + 1.0f) - sides.x) + (sides.x * sides.w));
		       area *= width;
		       if (width == 0.0f)
		       {
		           area = 0.0f;
		       }
		       fragCover.x = area;
		   }

		   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
		   {
		       vTo = stage_input.vTo;
		       vFrom = stage_input.vFrom;
		       vCtrl = stage_input.vCtrl;
		       frag_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.fragCover = fragCover;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
	shader_stencil_vert = backend.ShaderSources{
		Inputs: []backend.InputLocation{backend.InputLocation{Name: "corner", Location: 0, Semantic: "POSITION", SemanticIndex: 0, Type: 0x1, Size: 2}, backend.InputLocation{Name: "maxy", Location: 1, Semantic: "NORMAL", SemanticIndex: 0, Type: 0x0, Size: 1}, backend.InputLocation{Name: "from", Location: 2, Semantic: "TEXCOORD", SemanticIndex: 0, Type: 0x0, Size: 2}, backend.InputLocation{Name: "ctrl", Location: 3, Semantic: "TEXCOORD", SemanticIndex: 1, Type: 0x0, Size: 2}, backend.InputLocation{Name: "to", Location: 4, Semantic: "TEXCOORD", SemanticIndex: 2, Type: 0x0, Size: 2}},
		Uniforms: backend.UniformsReflection{
			Blocks:    []backend.UniformBlock{backend.UniformBlock{Name: "Block", Binding: 0}},
			Locations: []backend.UniformLocation{backend.UniformLocation{Name: "_16.transform", Type: 0x0, Size: 4, Offset: 0}, backend.UniformLocation{Name: "_16.pathOffset", Type: 0x0, Size: 2, Offset: 16}},
			Size:      24,
		},
		GLSL100ES: "#version 100\n\nstruct Block\n{\n    vec4 transform;\n    vec2 pathOffset;\n};\n\nuniform Block _16;\n\nattribute vec2 from;\nattribute vec2 ctrl;\nattribute vec2 to;\nattribute float maxy;\nattribute ivec2 corner;\nvarying vec2 vFrom;\nvarying vec2 vCtrl;\nvarying vec2 vTo;\n\nvoid main()\n{\n    vec2 from_1 = from + _16.pathOffset;\n    vec2 ctrl_1 = ctrl + _16.pathOffset;\n    vec2 to_1 = to + _16.pathOffset;\n    float maxy_1 = maxy + _16.pathOffset.y;\n    vec2 pos;\n    if (corner.x > 0)\n    {\n        pos.x = max(max(from_1.x, ctrl_1.x), to_1.x) + 1.0;\n    }\n    else\n    {\n        pos.x = min(min(from_1.x, ctrl_1.x), to_1.x) - 1.0;\n    }\n    if (corner.y > 0)\n    {\n        pos.y = maxy_1 + 1.0;\n    }\n    else\n    {\n        pos.y = min(min(from_1.y, ctrl_1.y), to_1.y) - 1.0;\n    }\n    vFrom = from_1 - pos;\n    vCtrl = ctrl_1 - pos;\n    vTo = to_1 - pos;\n    pos = (pos * _16.transform.xy) + _16.transform.zw;\n    gl_Position = vec4(pos, 1.0, 1.0);\n}\n\n",
		GLSL300ES: "#version 300 es\n\nlayout(std140) uniform Block\n{\n    vec4 transform;\n    vec2 pathOffset;\n} _16;\n\nlayout(location = 2) in vec2 from;\nlayout(location = 3) in vec2 ctrl;\nlayout(location = 4) in vec2 to;\nlayout(location = 1) in float maxy;\nlayout(location = 0) in ivec2 corner;\nout vec2 vFrom;\nout vec2 vCtrl;\nout vec2 vTo;\n\nvoid main()\n{\n    vec2 from_1 = from + _16.pathOffset;\n    vec2 ctrl_1 = ctrl + _16.pathOffset;\n    vec2 to_1 = to + _16.pathOffset;\n    float maxy_1 = maxy + _16.pathOffset.y;\n    vec2 pos;\n    if (corner.x > 0)\n    {\n        pos.x = max(max(from_1.x, ctrl_1.x), to_1.x) + 1.0;\n    }\n    else\n    {\n        pos.x = min(min(from_1.x, ctrl_1.x), to_1.x) - 1.0;\n    }\n    if (corner.y > 0)\n    {\n        pos.y = maxy_1 + 1.0;\n    }\n    else\n    {\n        pos.y = min(min(from_1.y, ctrl_1.y), to_1.y) - 1.0;\n    }\n    vFrom = from_1 - pos;\n    vCtrl = ctrl_1 - pos;\n    vTo = to_1 - pos;\n    pos = (pos * _16.transform.xy) + _16.transform.zw;\n    gl_Position = vec4(pos, 1.0, 1.0);\n}\n\n",
		/*
		   cbuffer Block : register(b0)
		   {
		       float4 _16_transform : packoffset(c0);
		       float2 _16_pathOffset : packoffset(c1);
		   };


		   static float4 gl_Position;
		   static float2 from;
		   static float2 ctrl;
		   static float2 to;
		   static float maxy;
		   static int2 corner;
		   static float2 vFrom;
		   static float2 vCtrl;
		   static float2 vTo;

		   struct SPIRV_Cross_Input
		   {
		       int2 corner : POSITION;
		       float maxy : NORMAL;
		       float2 from : TEXCOORD0;
		       float2 ctrl : TEXCOORD1;
		       float2 to : TEXCOORD2;
		   };

		   struct SPIRV_Cross_Output
		   {
		       float2 vFrom : TEXCOORD0;
		       float2 vCtrl : TEXCOORD1;
		       float2 vTo : TEXCOORD2;
		       float4 gl_Position : SV_Position;
		   };

		   void vert_main()
		   {
		       float2 from_1 = from + _16_pathOffset;
		       float2 ctrl_1 = ctrl + _16_pathOffset;
		       float2 to_1 = to + _16_pathOffset;
		       float maxy_1 = maxy + _16_pathOffset.y;
		       float2 pos;
		       if (corner.x > 0)
		       {
		           pos.x = max(max(from_1.x, ctrl_1.x), to_1.x) + 1.0f;
		       }
		       else
		       {
		           pos.x = min(min(from_1.x, ctrl_1.x), to_1.x) - 1.0f;
		       }
		       if (corner.y > 0)
		       {
		           pos.y = maxy_1 + 1.0f;
		       }
		       else
		       {
		           pos.y = min(min(from_1.y, ctrl_1.y), to_1.y) - 1.0f;
		       }
		       vFrom = from_1 - pos;
		       vCtrl = ctrl_1 - pos;
		       vTo = to_1 - pos;
		       pos = (pos * _16_transform.xy) + _16_transform.zw;
		       gl_Position = float4(pos, 1.0f, 1.0f);
		   }

		   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
		   {
		       from = stage_input.from;
		       ctrl = stage_input.ctrl;
		       to = stage_input.to;
		       maxy = stage_input.maxy;
		       corner = stage_input.corner;
		       vert_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.gl_Position = gl_Position;
		       stage_output.vFrom = vFrom;
		       stage_output.vCtrl = vCtrl;
		       stage_output.vTo = vTo;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
)
