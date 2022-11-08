<script>
    import Select, { Option } from '@smui/select';
    import Button, { Label } from '@smui/button';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import { getAllUser, connectGroup } from '../User/common-utils.js';
    import { onMount } from 'svelte';
    import { navigate } from "svelte-routing";
    
    export let id;

    let userId = '';
    let users = []

    function goToBack() {
        navigate("/usergroup/detail/" + id, {replace: true });
    }

    async function connectUser(uid, gid) {
        await connectGroup(uid, gid);
        navigate("/usergroup/detail/" + id, {replace: true });
    }

    onMount(async () => {
        users = await getAllUser();
    });
</script>

<CommonAppBar/>

<main>
    <div class="mdc-typography--headline3">Select User</div>
    <Select bind:value={userId} label="User" style="width: 70%">
        <Option value="" />
        {#each users as user}
          <Option value={user.id}>{user.name}</Option>
        {/each}
    </Select>
    <div class="inner">
        <div>
            <Button on:click={goToBack} variant="outlined" style="min-width: 170px">
                <Label>Cancel</Label>
            </Button>
        </div>
        <div>
            <Button on:click={connectUser(userId, id)} variant="unelevated" style="min-width: 170px">
                <Label>Save</Label>
            </Button>
        </div>
        
    </div>
</main>

<style>
    main {
		width: 100%;
		height: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    div.inner {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        margin-top: 20px;
    }
    div.inner > * {
        margin-right: 20px;
        margin-left: 20px;
    }
</style>