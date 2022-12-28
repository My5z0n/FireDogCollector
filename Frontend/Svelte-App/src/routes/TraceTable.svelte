<script>
	export let page;
	class Login {
		constructor(first, last,call) {
			this.first = first
			this.last = last
			this.call = call
		}
	}
	export let resultArray = [];
	export let GetTraceAddresses= 'http://127.0.0.1:9900/api/v1/traces/';
	export let GetTraceAddressesGet= '';
	export let testData = ""

	$: {
	GetTraceAddressesGet = GetTraceAddresses+'?page='+page;
	setTable();
	}

	function setTable(){
		console.log(GetTraceAddressesGet);
		const getRandomUser = async () => {
			fetch(GetTraceAddressesGet)
			.then(response => response.json())
			.then(data => {
				console.log(data);
				testData = data;
			}).catch(error => {
				 console.log(error);
				 return [];
				});
			}
		getRandomUser();
	
	}
</script>



<tbody>
	{#each testData as { TraceID, StartTime,Anomaly }}
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
