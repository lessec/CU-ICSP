<script>
	import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
    import IconButton, { Icon } from '@smui/icon-button';
	import Button, { Label } from '@smui/button';
	import CommonAppBar from '../Components/CommonAppBar.svelte';

	import { navigate } from "svelte-routing";
	import { onMount } from "svelte";
	import { Auth } from 'aws-amplify';

    let users = [];

	// get all users from server
    async function getAlls(e) {
		let response = await fetch("http://127.0.0.1:5555/api/v1/user/get/all");
		let commits = await response.json();
		users = commits;
	}

	// go to add user page
	function addNewUser() {
		navigate("user/add", {replace: true });
	}

	function showDetails(id) {
		navigate('user/detail/' + id, { replace: true });
	}

	onMount(async () => {
		const user = Auth.currentUserInfo().then(data => {
			console.log(data);
		});
		Auth.currentSession()
		.then(data => {
			let idToken = data.getIdToken();
			console.dir(idToken);
			let email = idToken.payload.email;
			console.log(email);
		})
		.catch(err => console.log(err));

		getAlls();
	});

	
</script>

<CommonAppBar isRefresh={true}/>

<main>
    <div class="inner">
		<div class="table-above">
			<div class="mdc-typography--headline3">User</div>
			<Button on:click={addNewUser} variant="unelevated" style="min-width: 170px">
				<Label>Add User</Label>
			</Button>
		</div>
		<DataTable stickyHeader table$aria-label="User list" style="width: 70%; height: 70%">
			<Head>
			  <Row>
				<Cell style="width: 100%;">Email</Cell>
				<Cell style="width: 100%;">Name</Cell>
				<Cell></Cell>
			  </Row>
			</Head>
			<Body>
			  {#each users as user}
				<Row>
				  <Cell>{user.email}</Cell>
				  <Cell>{user.name}</Cell>
				  <Cell>
					<div style="display: flex; align-items: center;">
						<IconButton class="material-icons" on:click={showDetails(user.id)} size="button">
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