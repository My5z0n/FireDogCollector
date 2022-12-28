<script>
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let traceID;
	export let SpanList = [];
	let GetTraceAddresses= 'http://127.0.0.1:9900/api/v1/traces/';
	let GetTraceAddressesGet= GetTraceAddresses+traceID;
	console.log("Gen");
	setTable();
	
	export let anomaly = false;

	function setSpanAttr(sID,anom,expected,got){
		dispatch('message', {
			spanID: sID,
			a: anom,
			e: expected,
			g: got,
		});
	}

	function setTable(){
		console.log(GetTraceAddressesGet);
		const getInfo = async () => {
			fetch(GetTraceAddressesGet)
			.then(response => response.json())
			.then(data => {
				console.log(data);
				let tmp = data.SpansList;
				tmp.forEach(x => {
					x.start_time = new Date(x.start_time).toISOString();
					x.end_time = new Date(x.end_time).toISOString();
					if (x.AnomalyDetected == true){
						anomaly = true;
					}
				 });
				 tmp.sort((a,b) => a.start_time >= b.start_time);
				 console.log(tmp);
				 SpanList = tmp;
			}).catch(error => {
				 console.log(error);
				 return [];
				});
			}
			getInfo();
	
	};
</script>



<tbody>
	{#each SpanList as {trace_id,span_id,parent_span_id,span_name,start_time,end_time,AnomalyDetected,ExpectedAnomalySpanName}}
	<tr class="{AnomalyDetected === true ? 'table-danger' : ''}">
	<td><button on:click={() =>{setSpanAttr(span_id,AnomalyDetected,ExpectedAnomalySpanName,span_name)}} class="btn btn-info btn-sm" data-bs-toggle="modal"
		data-bs-target="#myModal" >Browse</button></td>
	<td>{span_id}</td>


	{#if parent_span_id != ""}
	<td>{parent_span_id}</td>
	{:else}
	<td><span class="badge bg-dark">Root</span></td>
	{/if}
	<td>{span_name}</td>
	<td>{start_time}</td>
	<td>{end_time}</td>

	</tr>
	{/each}
  </tbody>

<style>

</style>
