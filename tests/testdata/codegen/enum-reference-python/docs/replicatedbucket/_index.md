
---
title: "ReplicatedBucket"
title_tag: "example.ReplicatedBucket"
meta_desc: "Documentation for the example.ReplicatedBucket resource with examples, input properties, output properties, lookup functions, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by test. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->




## Create ReplicatedBucket Resource {#create}
<div>
<pulumi-chooser type="language" options="typescript,python,go,csharp,java,yaml"></pulumi-chooser>
</div>


<div>
<pulumi-choosable type="language" values="javascript,typescript">
<div class="highlight"><pre class="chroma"><code class="language-typescript" data-lang="typescript"><span class="k">new </span><span class="nx">ReplicatedBucket</span><span class="p">(</span><span class="nx">name</span><span class="p">:</span> <span class="nx">string</span><span class="p">,</span> <span class="nx">args</span><span class="p">:</span> <span class="nx"><a href="#inputs">ReplicatedBucketArgs</a></span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#CustomResourceOptions">CustomResourceOptions</a></span><span class="p">);</span></code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"><span class=nd>@overload</span>
<span class="k">def </span><span class="nx">ReplicatedBucket</span><span class="p">(</span><span class="nx">resource_name</span><span class="p">:</span> <span class="nx">str</span><span class="p">,</span>
                     <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">Optional[ResourceOptions]</a></span> = None<span class="p">,</span>
                     <span class="nx">destination_region</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">)</span>
<span class=nd>@overload</span>
<span class="k">def </span><span class="nx">ReplicatedBucket</span><span class="p">(</span><span class="nx">resource_name</span><span class="p">:</span> <span class="nx">str</span><span class="p">,</span>
                     <span class="nx">args</span><span class="p">:</span> <span class="nx"><a href="#inputs">ReplicatedBucketArgs</a></span><span class="p">,</span>
                     <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">Optional[ResourceOptions]</a></span> = None<span class="p">)</span></code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"><span class="k">func </span><span class="nx">NewReplicatedBucket</span><span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">name</span><span class="p"> </span><span class="nx">string</span><span class="p">,</span> <span class="nx">args</span><span class="p"> </span><span class="nx"><a href="#inputs">ReplicatedBucketArgs</a></span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#ResourceOption">ResourceOption</a></span><span class="p">) (*<span class="nx">ReplicatedBucket</span>, error)</span></code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public </span><span class="nx">ReplicatedBucket</span><span class="p">(</span><span class="nx">string</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ReplicatedBucketArgs</a></span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.CustomResourceOptions.html">CustomResourceOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span></code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma">
<code class="language-java" data-lang="java"><span class="k">public </span><span class="nx">ReplicatedBucket</span><span class="p">(</span><span class="nx">String</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ReplicatedBucketArgs</a></span><span class="p"> </span><span class="nx">args<span class="p">)</span>
<span class="k">public </span><span class="nx">ReplicatedBucket</span><span class="p">(</span><span class="nx">String</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ReplicatedBucketArgs</a></span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">CustomResourceOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml">type: <span class="nx">example:ReplicatedBucket</span><span class="p"></span>
<span class="p">properties</span><span class="p">: </span><span class="c">#&nbsp;The arguments to resource properties.</span>
<span class="p"></span><span class="p">options</span><span class="p">: </span><span class="c">#&nbsp;Bag of options to control resource&#39;s behavior.</span>
<span class="p"></span>
</code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ReplicatedBucketArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#CustomResourceOptions">CustomResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>resource_name</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ReplicatedBucketArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">ResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">

<dl class="resources-properties"><dt
        class="property-optional" title="Optional">
        <span>ctx</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span>
    </dt>
    <dd>Context object for the current deployment.</dd><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ReplicatedBucketArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#ResourceOption">ResourceOption</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="csharp">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ReplicatedBucketArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.CustomResourceOptions.html">CustomResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ReplicatedBucketArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>options</span>
        <span class="property-indicator"></span>
        <span class="property-type">CustomResourceOptions</span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

## ReplicatedBucket Resource Properties {#properties}

To learn more about resource properties and how to use them, see [Inputs and Outputs](/docs/intro/concepts/inputs-outputs) in the Architecture and Concepts docs.

### Inputs

The ReplicatedBucket resource accepts the following [input](/docs/intro/concepts/inputs-outputs) properties:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="destinationregion_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#destinationregion_csharp" style="color: inherit; text-decoration: inherit;">Destination<wbr>Region</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Region to which data should be replicated.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="destinationregion_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#destinationregion_go" style="color: inherit; text-decoration: inherit;">Destination<wbr>Region</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Region to which data should be replicated.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="destinationregion_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#destinationregion_java" style="color: inherit; text-decoration: inherit;">destination<wbr>Region</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>Region to which data should be replicated.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="destinationregion_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#destinationregion_nodejs" style="color: inherit; text-decoration: inherit;">destination<wbr>Region</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Region to which data should be replicated.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="destination_region_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#destination_region_python" style="color: inherit; text-decoration: inherit;">destination_<wbr>region</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd><p>Region to which data should be replicated.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="destinationregion_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#destinationregion_yaml" style="color: inherit; text-decoration: inherit;">destination<wbr>Region</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>Region to which data should be replicated.</p>
</dd></dl>
</pulumi-choosable>
</div>


### Outputs

All [input](#inputs) properties are implicitly available as output properties. Additionally, the ReplicatedBucket resource produces the following output properties:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="locationpolicy_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_csharp" style="color: inherit; text-decoration: inherit;">Location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscaling">Pulumi.<wbr>Example.<wbr>Gcp/gke.<wbr>Outputs.<wbr>Node<wbr>Pool<wbr>Autoscaling</a></span>
    </dt>
    <dd><p>test stuff</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="locationpolicy_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_go" style="color: inherit; text-decoration: inherit;">Location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscaling">Node<wbr>Pool<wbr>Autoscaling</a></span>
    </dt>
    <dd><p>test stuff</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="locationpolicy_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_java" style="color: inherit; text-decoration: inherit;">location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscaling">Node<wbr>Pool<wbr>Autoscaling</a></span>
    </dt>
    <dd><p>test stuff</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="locationpolicy_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_nodejs" style="color: inherit; text-decoration: inherit;">location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscaling">gcpgke<wbr>Node<wbr>Pool<wbr>Autoscaling</a></span>
    </dt>
    <dd><p>test stuff</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="location_policy_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#location_policy_python" style="color: inherit; text-decoration: inherit;">location_<wbr>policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscaling">Node<wbr>Pool<wbr>Autoscaling</a></span>
    </dt>
    <dd><p>test stuff</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="locationpolicy_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_yaml" style="color: inherit; text-decoration: inherit;">location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscaling">Property Map</a></span>
    </dt>
    <dd><p>test stuff</p>
</dd></dl>
</pulumi-choosable>
</div>







## Supporting Types



<h4 id="nodepoolautoscaling">Node<wbr>Pool<wbr>Autoscaling</h4>

<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="locationpolicy_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_csharp" style="color: inherit; text-decoration: inherit;">Location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscalinglocationpolicy">Pulumi.<wbr>Google<wbr>Native.<wbr>Container/v1.<wbr>Node<wbr>Pool<wbr>Autoscaling<wbr>Location<wbr>Policy</a></span>
    </dt>
    <dd><p>Location policy used when scaling up a nodepool.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="locationpolicy_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_go" style="color: inherit; text-decoration: inherit;">Location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscalinglocationpolicy">Node<wbr>Pool<wbr>Autoscaling<wbr>Location<wbr>Policy</a></span>
    </dt>
    <dd><p>Location policy used when scaling up a nodepool.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="locationpolicy_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_java" style="color: inherit; text-decoration: inherit;">location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscalinglocationpolicy">Node<wbr>Pool<wbr>Autoscaling<wbr>Location<wbr>Policy</a></span>
    </dt>
    <dd><p>Location policy used when scaling up a nodepool.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="locationpolicy_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_nodejs" style="color: inherit; text-decoration: inherit;">location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscalinglocationpolicy">pulumi<wbr>Google<wbr>Nativecontainerv1Node<wbr>Pool<wbr>Autoscaling<wbr>Location<wbr>Policy</a></span>
    </dt>
    <dd><p>Location policy used when scaling up a nodepool.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="location_policy_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#location_policy_python" style="color: inherit; text-decoration: inherit;">location_<wbr>policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscalinglocationpolicy">Node<wbr>Pool<wbr>Autoscaling<wbr>Location<wbr>Policy</a></span>
    </dt>
    <dd><p>Location policy used when scaling up a nodepool.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="locationpolicy_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#locationpolicy_yaml" style="color: inherit; text-decoration: inherit;">location<wbr>Policy</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#nodepoolautoscalinglocationpolicy">&#34;LOCATION_POLICY_UNSPECIFIED&#34; | &#34;BALANCED&#34; | &#34;ANY&#34;</a></span>
    </dt>
    <dd><p>Location policy used when scaling up a nodepool.</p>
</dd></dl>
</pulumi-choosable>
</div>


<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="">example </a></dd>
	<dt>License</dt>
	<dd></dd>
</dl>

