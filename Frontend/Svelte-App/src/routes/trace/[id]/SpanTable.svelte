<script>
	export let traceID;
	let SpanList = [];
	let GetTraceAddresses= 'http://localhost:9900/api/v1/traces/';
	let GetTraceAddressesGet= GetTraceAddresses+traceID;
	console.log("Gen");
	setTable();
	



	function setTable(){
		console.log(GetTraceAddressesGet);
		const getInfo = async () => {
			fetch(GetTraceAddressesGet)
			.then(response => response.json())
			.then(data => {
				console.log(data);
				SpanList = data.SpansList;
			}).catch(error => {
				 console.log(error);
				 return [];
				});
			}
			getInfo();
	
	}
</script>



<tbody>
	{#each SpanList as {trace_id,span_id,parent_span_id,span_name,start_time,end_time,AnomalyDetected,ExpectedAnomalySpanName}}
	<tr>
	<td><span class="badge btn bg-info" data-bs-toggle="modal"
		data-bs-target="#myModal">Get More</span></td>
	<td>{span_id}</td>
	<td>{parent_span_id}</td>
	<td>{span_name}</td>
	<td>{start_time}</td>
	<td>{new Date(end_time).toISOString()}</td>
	<td>{AnomalyDetected}</td>
	</tr>
	{/each}
  </tbody>

<style>

</style>
