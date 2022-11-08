<script>
	import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
    import IconButton, { Icon } from '@smui/icon-button';
	import Button, { Label } from '@smui/button';
	import CommonAppBar from '../Components/CommonAppBar.svelte';

	import { onMount } from "svelte";
	import { navigate } from "svelte-routing";

    let groups = [];

	// get all groups from server
    async function getAlls(e) {
		let response = await fetch("http://127.0.0.1:5555/api/v1/group/get/all");
		let commits = await response.json();
		groups = commits;
	}

	// go to add group page
	function addNewUserGroup() {
		navigate("usergroup/add", {replace: true });
	}

	onMount(async () => {
		await getAlls();
	});

	// go to edit group page
    function editUserGroup(id) {
		navigate("usergroup/edit/" + id, {replace: true });
	}

	function showDetails(id) {
		navigate("usergroup/detail/" + id, {replace: true });
	}

</script>

<CommonAppBar isRefresh={true} />

<main>
    <div class="inner">
		<div class="table-above">
			<div class="mdc-typography--headline3">Group</div>
			<Button on:click={addNewUserGroup} variant="unelevated" style="min-width: 170px">
				<Label>Add Group</Label>
			</Button>
		</div>
		<DataTable stickyHeader table$aria-label="User list" style="width: 70%; height: 70%">
			<Head>
			  <Row>
				<Cell style="width: 100%;">Name</Cell>
				<Cell></Cell>
			  </Row>
			</Head>
			<Body>
			  {#each groups as group}
				<Row>
				  <Cell>{group.name}</Cell>
				  <Cell>
					<div style="display: flex; align-items: center;">
						<IconButton class="material-icons" on:click={showDetails(group.id)} size="button">
							more_horiz
						</IconButton>
					  </div>
				  </Cell>
				</Row>
			  {/each}
			</Body>
		</DataTable>
    </div>
</main>

<style>
    main {
		width: 100%;
		height: 100%;
        display: flex;
        justify-content: center;
		margin: 20px;
    }
    div.inner {
		width: 100%;
		height: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
	div.table-above {
		display: flex;
		flex-direction: row;
		margin-bottom: 20px;
		justify-content: space-between;
		align-items: center;
		width: 70%;
	}
</style>