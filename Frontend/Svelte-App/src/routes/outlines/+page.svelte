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
    console.log("fetching data");
    console.log(JSON.stringify(anomalyElementObj));
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
          console.log(data);
          result = data;
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
  <h1>Root cause analys results:</h1>
</div>

<div class="bg-light p-4 rounded">
  <table>
    <thead>
      <tr>
        <th>Column</th>
        <th>Value</th>
      </tr>
    </thead>
    <tbody>
      {#if result != null}
        {#each result[0].itemSets[0].items as item}
          <tr>
            <td>{item.column}</td>
            <td>{item.value ?? "null"}</td>
          </tr>
        {/each}
      {/if}
    </tbody>
  </table>
</div>
