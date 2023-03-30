<script>
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import { anomalyElementObjStore } from "../../store.js";

  let anomalyElementObj;
  let result = null;
  let x = null;

  $: {
    anomalyElementObj = $anomalyElementObjStore;
  }
  let SpanObj;

  function fetchData() {
    //console.log("fetching data");
    //console.log(JSON.stringify(anomalyElementObj));
    const get = async () => {
      fetch("http://127.0.0.1:9900/api/v1/anomalydetetor/find-outlines", {
        method: "POST",
        body: JSON.stringify(anomalyElementObj),
        headers: {
          "Content-type": "application/json",
        },
      })
        .then((response) => response.json())
        .then((data) => {
          result = data.map((result) => ({
            ...result,
            itemSets: result.itemSets.map((itemSet) => ({
              ...itemSet,
              items: itemSet.items.filter((item) => item.column !== "Test"),
            })),
          }));
        })
        .catch((error) => {
          console.log(error);
          return [];
        });
    };
    get();
  }
  onMount(() => {
    SpanObj = $page.myValue;
    fetchData();
  });
</script>

<svelte:head>
  <title>Outline Trace</title>
  <meta name="description" content="Trace this app" />
</svelte:head>

<div class="bg-light p-4 rounded">
  <h1 class="h4">Root cause analysis results:</h1>
</div>

<div class="bg-light p-4 rounded">
  <div class="row">
    <div class="col-12">
      {#if result != null}
        <h6 class="mb-3">Num Outliners: {result[0].numOutliers}</h6>
        <h6 class="mb-3">Num Inliers: {result[0].numInliers}</h6>
        <h6 class="mb-3">Execution Time: {result[0].executionTime}</h6>
        <h6 class="mb-3">Load Time: {result[0].loadTime}</h6>
        <h6 class="mb-3">Summarization Time: {result[0].summarizationTime}</h6>

        {#each result[0].itemSets as itemSet}
          <hr />
          <div class="mb-5">
            <h6 class="mb-3">Support: {itemSet.support}</h6>
            <h6 class="mb-3">Num Records: {itemSet.numRecords}</h6>
            <h6 class="mb-3">Ratio to Inliers: {itemSet.ratioToInliers}</h6>
            <table class="table table-bordered">
              <thead>
                <tr>
                  <th>Column</th>
                  <th>Value</th>
                </tr>
              </thead>
              <tbody>
                {#each itemSet.items as item}
                  <tr>
                    <td>{item.column}</td>
                    <td>{item.value ?? "null"}</td>
                  </tr>
                {/each}
              </tbody>
            </table>
          </div>
        {/each}
      {/if}
    </div>
  </div>
</div>
