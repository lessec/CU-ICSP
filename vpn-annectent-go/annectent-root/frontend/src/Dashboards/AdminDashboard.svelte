<script>
    import Switch from '@smui/switch';
    import FormField from '@smui/form-field';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    let isWorking = false;

    let host = 'http://localhost:5555';
    let api = '/api/v1/wiregaurd'
        
    async function wireguardSwitch(e) {
        if(e['detail']['selected']) {
            console.log('on');
            let url = host + api + '/up';
            await fetch(url);
        } else {
            console.log('off');
            let url = host + api + '/down';
            await fetch(url);
        }
    }
</script>

<CommonAppBar isRefresh={true}/>

<main>
    <div class="topline">
        <div class="mdc-typography--headline4">Annectect status</div>
        <FormField>
            <Switch bind:checked={isWorking} on:SMUISwitch:change={(e)=>{wireguardSwitch(e)}} />
            <span slot="label">{isWorking?"Working now!":"Idle"}</span>
        </FormField>
    </div>
</main>