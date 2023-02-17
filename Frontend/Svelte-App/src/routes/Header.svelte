<script>

let GetTraceAddresses= 'http://127.0.0.1:9900/api/v1/traces/';

function CallAPITrainModel(){
	const get = async () => {
		fetch(GetTraceAddresses)
		.then(response => response.json())
		.catch(error => {
			 console.log(error);
			 return [];
			});
		}
		get();
}
</script>

<header>
<!-- Image and text -->
<div>
	<nav class="navbar navbar-expand navbar-light bg-light px-5">
		<a class="navbar-brand" href="/">🐺 Firedog </a>
		  <div class="navbar-nav ">
			<!-- 			<a class=" nav-item nav-link active" href="#">Trace</a> -->

			<a on:click={CallAPITrainModel} class="nav-item nav-link" href="#">Train Model</a>
		</div>
</div>


</header>

<style>

</style>
