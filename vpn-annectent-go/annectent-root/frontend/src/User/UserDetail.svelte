<script>
    import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
    import IconButton, { Icon } from '@smui/icon-button';
	import Button, { Label } from '@smui/button';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import { onMount } from 'svelte';
    import { navigate } from "svelte-routing";

    export let id;

    let devices = [];
    let name = '';
    let email = '';
    let type = '';
    let group = '';

    function addDevice() {
        navigate('/user/detail/' + id + '/adddevice', {replace: true});
    }

    // get an user from server
    async function getUser(id) {
        let url = `http://127.0.0.1:5555/api/v1/user/get/id/${id}`;
        
        let response = await fetch(url);
        console.log(response);
        let data = await response.json();
        console.log(data);
        return data;
    }

    function showDeviceDetail(deviceId) {
        navigate('/user/detail/' + id + '/device/' + deviceId, {replace: true});
    }

    function goToBack() {
        navigate('/user', { replace: true });
    }

    onMount(async () => {
        let user = await getUser(id);
        console.log(user);
        name = user.name;
        email = user.email;
        type = user.type;
        group = user.groups == null ? '' : user.groups[0].Name;
        devices = user.devices == null ? [] : user.devices;
    });

    function goToEdit() {
        navigate('/user/edit/' + id, { replace: true });
    }
</script>

<CommonAppBar isSetting={true} action={goToEdit}/>

<main>
    <div class="top-line">
        <div class="button" style="display: flex; align-items: center;">
            <IconButton class="material-icons" on:click={goToBack}>
                arrow_back
            </IconButton>
        </div>
        <div class="mdc-typography--headline2">{name}</div>
        <div class="mdc-typography--subtitle1 ml">({type})</div>
    </div>
    <div class="mid-line">
        <div class="mdc-typography--headline3">{group}</div>
        <div class="mdc-typography--body1 mr">{email}</div>
    </div>
    <div class="bot-line">
        <div class="add-button">
            <Button on:click={addDevice(id)} variant="unelevated" style="min-width: 170px">
                <Label>Add device</Label>
            </Button>
        </div>
        
        <DataTable stickyHeader table$aria-label="User list" style="width: 100%; height: 70%">
			<Head>
			  <Row>
                <Cell style="width: 100%;">Name</Cell>
				<Cell></Cell>
			  </Row>
			</Head>
			<Body>
			  {#each devices as device}
				<Row>
				  <Cell>{device.Name}</Cell>
				  <Cell>
					<div style="display: flex; align-items: center;">
						<IconButton class="material-icons" on:click={showDeviceDetail(device.ID)} size="button">
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
        align-items: center;
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
        margin-right: 30px;
    }
    .ml {
        margin-left: 10px;
    }
</style>