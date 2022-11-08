<script>
    import IconButton, { Icon } from '@smui/icon-button';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import {navigate} from "svelte-routing";
    import {onMount} from "svelte";

    export let id;
    export let did;

    let name = '';
    let ip = '';
    let publicKey = '';
    let imgUrl = `http://127.0.0.1:5555/api/v1/device/get/id/${did}?format=qr&a=${new Date().getTime()}`;

    // get a device from server
    async function getData(did) {
        let url = `http://127.0.0.1:5555/api/v1/device/get/id/${did}`

        let response = await fetch(url);
        console.log(response);
        return await response.json();
    }

    // lifecycle function
    onMount(async () => {
        getData(did).then(data => {
            name = data.name;
            ip = data.ip;
            publicKey = data.publicKey;
        });
    });

    function goToBack() {
        navigate('/user/detail/' + id, {replace: true});
    }

    function goToEdit() {
        navigate('/user/detail/' + id + '/editdevice/' + did, {replace: true});
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
        <div class="address">
            <div class="mdc-typography--headline4">Address</div>
            <div class="mdc-typography--body1">{ip}</div>
        </div>
        <div class="pubkey">
            <div class="mdc-typography--headline4">Public Key</div>
            <div class="mdc-typography--body1">{publicKey}</div>
        </div>
    </div>
    <div class="bot-line">
        <img src="{imgUrl}" id="qr-code" alt="Mobile client config"/>
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
        flex-direction: column;
        min-height: 180px;
    }
    div.mid-line > div {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        width: 100%;
    }
    div.bot-line {
        display: flex;
        flex-direction: column;
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
</style>