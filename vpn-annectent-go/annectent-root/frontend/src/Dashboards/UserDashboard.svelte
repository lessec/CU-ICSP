<script>
    import {onMount} from 'svelte';
    import {UserStore} from "../Stores/user-store.js";
    import {getAllGroups} from "../UserGroup/common-utils.js";
    import {getAllDevices} from "../UserDevice/common-utils.js";
    import {connectGroup, disconnectGroup, connectDevice, disconnectDevice} from "../User/common-utils.js";

    let username = '';
    let name = '';
    let email = '';
    let type = '';
    let id = '';
    let groups = [];
    let devices = [];
    let groupId = -1;
    let deviceId = -1;
    let allGroups = [];
    let allDevices = [];
   
    async function getUserInfo(){
        if ($UserStore.username.startsWith('loginwithamazon')
            ||$UserStore.username.startsWith('google_')) {
            username = $UserStore.attributes.name;
        } else {
            username = $UserStore.username;
        }

        let res = await fetch(`http://127.0.0.1:5555/api/v1/user/get/u/${username}`)
        let data = await res.json();
        console.log(data);
        name = data.name;
        email = data.email;
        type = data.type;
        id = data.id;
        groups = data.groups==null?[]:data.groups;
        devices = data.devices==null?[]:data.devices;

        allGroups = await getAllGroups();
        allDevices = await getAllDevices();
    }

    onMount(async () => {
        await getUserInfo();
    })
    
</script>

<h1>{name}</h1>
<h3>{email}</h3>
<h3>{type}</h3>

<!-- group add -->
<select bind:value={groupId}>
    <option value={-1}>Select Group</option>
    {#each allGroups as group}
        <option value={group.id}>{group.name}</option>
    {/each}
</select>
{#if groupId!=-1}
    <button on:click={async ()=>{
        await connectGroup(id, groupId);
        await getUserInfo();
    }}>Add Group To User</button>
{/if}

<!-- device add -->
<select bind:value={deviceId}>
    <option value={-1}>Select Device</option>
    {#each allDevices as device}
        <option value={device.id}>{device.name}</option>
    {/each}
</select>
{#if deviceId!=-1}
    <button on:click={async ()=>{
        await connectDevice(id, deviceId);
        await getUserInfo();
    }}>Add Device to User</button>
{/if}

<!-- show groups related -->
{#each groups as group}
        <h3>{group.Name}</h3>
        <button on:click={async ()=>{
            await disconnectGroup(id, group.ID);
            await getUserInfo();
        }}>x</button>
{/each}

<!-- show devices related -->
{#each devices as device}
        <h3>{device.Name}</h3>
        <button on:click={async ()=>{
            await disconnectDevice(id, device.ID);
            await getUserInfo();
        }}>x</button>
{/each}

