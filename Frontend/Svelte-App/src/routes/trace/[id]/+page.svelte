<svelte:head>
	<title>Trace</title>
	<meta name="description" content="Trace this app" />
</svelte:head>

<script>
	import SpanTable from './SpanTable.svelte';
	import Graph from './Graph.svelte';
	import { page } from '$app/stores';

	let name = $page.params.id;
	let anomaly = false;
	let SpanList = [];
	let modal_spanID='';
	let attr_spanID = [];
	let modal_expected = '';
	let modal_got = '';
	let modal_anomaly = false;

	

	function handleMessage(event) {
		modal_spanID = event.detail.spanID;
		modal_expected = event.detail.e;
		modal_anomaly = event.detail.a;
		modal_got = event.detail.g;
		setModal();

	}

	function setModal(){
		const get = async () => {
			fetch("http://127.0.0.1:9900/api/v1/spans/"+modal_spanID)
			.then(response => response.json())
			.then(data => {
				let tmp = data.Attributes;
				attr_spanID = Object.entries(tmp);
				attr_spanID.sort((a,b) => a[0].localeCompare(b[0]));
				console.log(attr_spanID); 
			}).catch(error => {
				 console.log(error);
				 return [];
				});
			}
			get();
	
	};
</script>

<div class="bg-light p-4 rounded">
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center  pb-2 mb-3 border-bottom">
		<h1 class="h2">Trace ID : [{name}] </h1>
		{#if anomaly}
		<div class="mb-2 mb-md-0">
			<h2><span class="badge bg-danger">Anomaly</span></h2>
		</div>
		{/if}
	</div>
	<Graph bind:SpanList={SpanList}/>
	<table class="table mt-2">
		<thead>
		  <tr >
			<th scope="col">Attributes</th>
			<th scope="col">SpanID</th>
			<th scope="col">Parent Span ID</th>
			<th scope="col">SpanName</th>
			<th scope="col">Start Time</th>
			<th scope="col">End Time</th>

		  </tr>
		</thead>
		<SpanTable bind:anomaly={anomaly} bind:SpanList={SpanList} traceID={name} on:message={handleMessage}/>
	  </table>


</div>

<!-- The Modal -->
<div class="modal" id="myModal">
	<div class="modal-dialog ">
		<div class="modal-content ">
			<div class="modal-header">
				<h1 class="modal-title fs-5" id="exampleModalLabel">
					Span attributes [{modal_spanID}] 
				</h1>
				<button
					type="button"
					class="btn-close"
					data-bs-dismiss="modal"
					aria-label="Close"
				/>
			</div>
			<div class="modal-body">
				{#if modal_anomaly}
				<h5 class="text-danger">Received: {modal_got}</h5>
				<h5 class="text-danger">Expected: {modal_expected}</h5>
				{/if}
				<table class="table">
					<thead>
					  <tr>
						<th scope="col">Attribute</th>
						<th scope="col">Value</th>
					  </tr>
					</thead>
					<tbody>
						{#each attr_spanID as [key, value]}
							<tr>
							<td>{key}</td>
							<td>{value}</td>
						  </tr>
						{/each}
		
					</tbody>
				  </table>
	
		</div>
		</div>
	</div>
</div>
