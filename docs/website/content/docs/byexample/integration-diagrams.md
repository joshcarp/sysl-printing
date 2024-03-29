+++
# AUTOGENERATED BY byexample/generate.go
title= "Integration Diagrams"
draft= false
description= ""
layout= "byexample"
weight = 10
topic = "Diagrams"
PlaygroundURL = "http://anz-bank.github.io/sysl-playground/?input=SW50ZWdyYXRlZFN5c3RlbToKICAgIGludGVncmF0ZWRfZW5kcG9pbnRfMToKICAgICAgICBTeXN0ZW0xIDwtIGVuZHBvaW50CiAgICBpbnRlZ3JhdGVkX2VuZHBvaW50XzI6CiAgICAgICAgU3lzdGVtMiA8LSBlbmRwb2ludApTeXN0ZW0xOgogICAgZW5kcG9pbnQ6IC4uLgpTeXN0ZW0yOgogICAgZW5kcG9pbnQ6IC4uLgpQcm9qZWN0IFthcHBmbXQ9IiUoYXBwbmFtZSkiXToKICAgIF86CiAgICAgICAgSW50ZWdyYXRlZFN5c3RlbQogICAgICAgIFN5c3RlbTEKICAgICAgICBTeXN0ZW0yCg==&cmd=c3lzbCBpbnRzIC1vIDNfcHJvamVjdC5zdmcgLS1wcm9qZWN0IFByb2plY3QgMV9wcm9qZWN0LnN5c2w="
ID = "integration-diagrams"
CodeWithoutComments = """IntegratedSystem:
    integrated_endpoint_1:
        System1 <- endpoint
    integrated_endpoint_2:
        System2 <- endpoint
System1:
    endpoint: ...
System2:
    endpoint: ...
Project [appfmt="%(appname)"]:
    _:
        IntegratedSystem
        System1
        System2
"""

Segs = [[
  
      {CodeEmpty= true,CodeLeading= true,CodeRun= false,CodeRendered="""""",DocsRendered= """<p>In this example will use a simple system and start using the sysl command to generate diagrams.</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= true,CodeRendered="""<pre class="chroma"><span class="nx">IntegratedSystem</span><span class="p">:</span>
    <span class="nx">integrated_endpoint_1</span><span class="p">:</span>
        <span class="nx">System1</span> <span class="o">&lt;-</span> <span class="nx">endpoint</span>
    <span class="nx">integrated_endpoint_2</span><span class="p">:</span>
        <span class="nx">System2</span> <span class="o">&lt;-</span> <span class="nx">endpoint</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma"><span class="nx">System1</span><span class="p">:</span>
    <span class="nx">endpoint</span><span class="p">:</span> <span class="o">...</span>
<span class="nx">System2</span><span class="p">:</span>
    <span class="nx">endpoint</span><span class="p">:</span> <span class="o">...</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= false,CodeRun= false,CodeRendered="""<pre class="chroma"><span class="nx">Project</span> <span class="p">[</span><span class="nx">appfmt</span><span class="p">=</span><span class="s">&#34;%(appname)&#34;</span><span class="p">]:</span>
    <span class="nx">_</span><span class="p">:</span>
        <span class="nx">IntegratedSystem</span>
        <span class="nx">System1</span>
        <span class="nx">System2</span></pre>""",DocsRendered= """""",Image = ""},


],
[
  
      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">
<span class="nx">export</span> <span class="nx">SYSL_PLANTUML</span><span class="p">=</span><span class="nx">http</span><span class="p">:</span><span class="o">//</span><span class="nx">www</span><span class="p">.</span><span class="nx">plantuml</span><span class="p">.</span><span class="nx">com</span><span class="o">/</span><span class="nx">plantuml</span></pre>""",DocsRendered= """<p>First, make sure to set the environment variable SYSL_PLANTUML</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= true,CodeRendered="""<pre class="chroma">
<span class="nx">sysl</span> <span class="nx">ints</span> <span class="o">-</span><span class="nx">o</span> <span class="mi">3</span><span class="nx">_project</span><span class="p">.</span><span class="nx">svg</span> <span class="o">--</span><span class="nx">project</span> <span class="nx">Project</span> <span class="mi">1</span><span class="nx">_project</span><span class="p">.</span><span class="nx">sysl</span></pre>""",DocsRendered= """<p>Now run the sysl sd (sequence diagram) command</p>
""",Image = ""},

      {CodeEmpty= true,CodeLeading= true,CodeRun= false,CodeRendered="""""",DocsRendered= """<p><code>-o</code> is the output file</p>
""",Image = ""},

      {CodeEmpty= true,CodeLeading= true,CodeRun= false,CodeRendered="""""",DocsRendered= """<p><code>-s</code> specifies a starting endpoint for the sequence diagram to initiate</p>
""",Image = ""},

      {CodeEmpty= true,CodeLeading= true,CodeRun= false,CodeRendered="""""",DocsRendered= """<p><code>project.sysl</code> is the input sysl file</p>
""",Image = ""},

      {CodeEmpty= true,CodeLeading= false,CodeRun= false,CodeRendered="""""",DocsRendered= """<p>project.svg:</p>
""",Image = ""},


],
[
  
      {CodeEmpty= false,CodeLeading= false,CodeRun= false,CodeRendered="""""",DocsRendered= """""",Image = "/assets/byexample/images/integration-diagrams9.svg"},


],

]
+++


