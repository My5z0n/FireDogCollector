<script>
	import SpanTable from "./SpanTable.svelte";
	import Graph from "./Graph.svelte";
	import { page } from "$app/stores";

	let name = $page.params.id;
	let DetectedAnomalyInTrace = false;
	let SpanList = [];
	let Modal_SpanID = "";
	let Modal_SpanID_Attributes = [];
	let Anomaly_SpanID = "";
	let ExpectedAnomalySpanName = "";
	let ReceivedAnomalySpanName = "";


	//Handle Modal See Span Atr List
	function handleMessage(event) {
		Modal_SpanID = event.detail.spanID;
		setModal();
	}

	function setModal() {
		const get = async () => {
			fetch("http://127.0.0.1:9900/api/v1/spans/" + Modal_SpanID)
				.then((response) => response.json())
				.then((data) => {
					let tmp = data.Attributes;
					Modal_SpanID_Attributes = Object.entries(tmp);
					Modal_SpanID_Attributes.sort((a, b) => a[0].localeCompare(b[0]));
				})
				.catch((error) => {
					console.log(error);
					return [];
				});
		};
		get();
	}
</script>

<svelte:head>
	<title>Trace</title>
	<meta name="description" content="Trace this app" />
</svelte:head>

<div class="bg-light p-4 rounded">
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center  pb-2 mb-3 border-bottom"
	>
		<h1 class="h2">Trace ID : [{name}]</h1>
		{#if DetectedAnomalyInTrace}
			<div class="mb-2 mb-md-0">
				<h2><span class="badge bg-danger">Anomaly</span></h2>
			</div>
		{/if}
	</div>
	<Graph bind:SpanList />

	{#if DetectedAnomalyInTrace}
		<hr />
		{#if Anomaly_SpanID!=""}
		<h5 class="text-danger">
			Detected anomaly in this Trace. Span ID: [{Anomaly_SpanID}]
		</h5>
		{/if}
		<h5 class="text-danger">Received: {ReceivedAnomalySpanName}</h5>
		<h5 class="text-danger">Expected: {ExpectedAnomalySpanName}</h5>
		<hr />
	{/if}

	<table class="table mt-2">
		<thead>
			<tr>
				<th scope="col">Attributes</th>
				<th scope="col">SpanID</th>
				<th scope="col">Parent Span ID</th>
				<th scope="col">SpanName</th>
				<th scope="col">Start Time</th>
				<th scope="col">End Time</th>
			</tr>
		</thead>
		<SpanTable
			bind:anomaly={DetectedAnomalyInTrace}
			bind:anomaly_SpanID={Anomaly_SpanID}
			bind:expected={ExpectedAnomalySpanName}
			bind:got={ReceivedAnomalySpanName}
			bind:SpanList
			traceID={name}
			on:message={handleMessage}
		/>
	</table>
</div>

<!-- The Modal -->
<div class="modal" id="myModal">
	<div class="modal-dialog ">
		<div class="modal-content ">
			<div class="modal-header">
				<h1 class="modal-title fs-5" id="exampleModalLabel">
					Span attributes [{Modal_SpanID}]
				</h1>
				<button
					type="button"
					class="btn-close"
					data-bs-dismiss="modal"
					aria-label="Close"
				/>
			</div>
			<div class="modal-body">
				<table class="table">
					<thead>
						<tr>
							<th scope="col">Attribute</th>
							<th scope="col">Value</th>
						</tr>
					</thead>
					<tbody>
						{#each Modal_SpanID_Attributes as [key, value]}
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
