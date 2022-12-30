<script>
	export let page;
	export let GetTraceAddresses= 'http://127.0.0.1:9900/api/v1/traces/';
	export let GetTraceAddressesGet= '';
	export let TracesList = ""

	$: {
	GetTraceAddressesGet = GetTraceAddresses+'?page='+page;
	setTable();
	}

	function setTable(){
		const getRandomUser = async () => {
			fetch(GetTraceAddressesGet)
			.then(response => response.json())
			.then(data => {
				TracesList = data;
			}).catch(error => {
				 console.log(error);
				 return [];
				});
			}
		getRandomUser();
	
	}
</script>



<tbody>
	{#each TracesList as { TraceID, StartTime,Anomaly }}
	<tr>
	<td><a type="button" class="btn btn-info btn-sm" href={"/trace/"+TraceID}>More Info</a></td>
	  <td>{TraceID}</td>
	  <td>{new Date(StartTime).toISOString()}</td>
	  {#if Anomaly.Valid}
		{#if Anomaly.Bool}
		<td><span class="badge bg-danger">Detected</span></td>
		{:else}
		<td><span class="badge bg-success">Clear</span></td>
		{/if}
		{:else}
		<td><span class="badge bg-secondary">Not Checked</span></td>
	  {/if}

	</tr>
	{/each}
  </tbody>

<style>

</style>
