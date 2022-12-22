<svelte:head>
	<title>Trace</title>
	<meta name="description" content="Trace this app" />
</svelte:head>

<script>
	import SpanTable from './SpanTable.svelte';
	import { page } from '$app/stores';
	let name = $page.params.id;

</script>

<div class="bg-light p-4 rounded">
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center  pb-2 mb-3 border-bottom">
		<h1 class="h2">Trace ID : [{name}] </h1>
	</div>
	<table class="table mt-1 w-auto">
		<thead>
		  <tr >
			<th scope="col">Info</th>
			<th scope="col">SpanID</th>
			<th scope="col">Parent Span ID</th>
			<th scope="col">SpanName</th>
			<th scope="col">Start Time</th>
			<th scope="col">End Time</th>
			<th scope="col">Anomaly</th>
		  </tr>
		</thead>
		<SpanTable traceID={name}/>
	  </table>


</div>

<!-- The Modal -->
<div class="modal" id="myModal">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<h1 class="modal-title fs-5" id="exampleModalLabel">
					Modal title
				</h1>
				<button
					type="button"
					class="btn-close"
					data-bs-dismiss="modal"
					aria-label="Close"
				/>
			</div>
			<div class="modal-body">...</div>
		</div>
	</div>
</div>
