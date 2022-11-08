<script>
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
    import IconButton, { Icon } from '@smui/icon-button';
	import Button, { Label } from '@smui/button';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import { onMount } from 'svelte';
    import { disconnectGroup } from '../User/common-utils.js';
    import { navigate } from "svelte-routing";

    export let id;

    let users = [];
    let name = '';
    let describe = '';

    async function getAllRelatedUsers() {
        let response = await fetch("http://127.0.0.1:5555/api/v1/group/users/" + id);
		let commits = await response.json();
		users = commits;
    }

    function addUser(id) {
        let url = "/usergroup/detail/" + id + "/adduser";
        navigate(url, { replace: true });
    }

    async function deleteUser(groupId, userId) {
        let answer = confirm("Are you sure you want to delete this user?");
        if (answer) {
            await disconnectGroup(userId, groupId);
            await getAllRelatedUsers();
        }
    }

    function goToBack() {
        navigate("/usergroup", { replace: true });
    }

    // get a group from server
    async function getUserGroup(id) {
        let url = `http://127.0.0.1:5555/api/v1/group/get/id/${id}`;
       
		let response = await fetch(url);
        console.log(response);
        let data = await response.json();
        return data;
	}

    onMount(async () => {
        await getAllRelatedUsers();
        let group = await getUserGroup(id);
        name = group.name;
        describe = group.describe;
    });

    function goToEdit() {
        navigate('/usergroup/edit/' + id, { replace: true });
    }
</script>

<CommonAppBar isSetting={true} action={goToEdit}/>

<main>
    <div class="top-line">
        <div class="button" style="display: flex; align-items: center;">
            <IconButton class="material-icons" on:click={goToBack}
              >arrow_back</IconButton>
        </div>
        <div class="mdc-typography--headline2">{name}</div>
    </div>
    <div class="mid-line">
        <div class="mdc-typography--headline3">Describe</div>
        <div class="mdc-typography--body1 mr">{describe}</div>
    </div>
    <div class="bot-line">
        <div class="add-button">
            <Button on:click={addUser(id)} variant="unelevated" style="min-width: 170px">
                <Label>Add User</Label>
            </Button>
        </div>
        
        <DataTable stickyHeader table$aria-label="User list" style="width: 100%; height: 70%">
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
				  <Cell>{user.name}</Cell>
                  <Cell>{user.email}</Cell>
				  <Cell>
					<div style="display: flex; align-items: center;">
						<IconButton class="material-icons" on:click={deleteUser(id, user.id)} size="button">
							close
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
        align-items: center;
        flex-direction: column;
		margin: 20px;
    }
    main > div {
		width: 70%;
		height: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    div.mid-line {
        display: flex;
        flex-direction: row;
        align-items: flex-start;
        justify-content: space-between;
        min-height: 180px;
    }
    div.bot-line {
        display: flex;
        flex-direction: column;
    }
    div.add-button {
        margin-top: 10px;
        margin-bottom: 20px;
    }
    div.top-line {
        display: flex;
        flex-direction: row;
        align-items: center;
        min-height: 60px;
        justify-content: center;
        margin-bottom: 10px;
    }
    div.top-line > div.button {
        position: absolute;
        left: 17%;
    }
    .mr {
        margin-left: 30px;
    }
</style>
