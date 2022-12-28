<script>
	import TraceTable from './TraceTable.svelte';

    let page_no = 0;

    function NextClick() {
        page_no += 1;
    }
    function PreviousClick() {
        page_no -= 1;
    }
</script>

<h1  class="">Traces from Tool:</h1>
<table class="table mt-1 w-auto">
    <thead>
      <tr >
        <th scope="col">Info</th>
        <th scope="col">TraceID</th>
        <th scope="col">StartTime</th>
        <th scope="col">Anomaly</th>
      </tr>
    </thead>
    <TraceTable page={page_no}/>
  </table>

  <nav aria-label="Page navigation example">
    <ul class="pagination justify-content-center">
      {#if page_no>0}
      <li class="page-item"><a  on:click={PreviousClick} class="page-link" href="#">Previous</a></li>
      {/if}
      <li class="page-item" ><span class="page-link" >{page_no}</span></li>
      <li class="page-item" ><a on:click={NextClick} class="page-link" href="#">Next</a></li>
    </ul>
  </nav>

  <style>

 
  </style>