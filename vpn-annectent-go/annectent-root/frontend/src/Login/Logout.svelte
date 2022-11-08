<script>
import { navigate } from "svelte-routing";
import Button, { Label } from '@smui/button';
import Logo from '../Assets/Logo.svelte';
</script>

<main>
    <div class="inner">
        <Logo />
        <div class="mdc-typography--headline2 mt">Good Bye</div>
        <div class="mdc-typography--subtitle1">You can leave our service.</div>
        <Button color="secondary" on:click={()=>{navigate('/', {replace: true})}} variant="unelevated" style="min-width: 300px; margin-top: 10px">
            <Label>Go to Back</Label>
        </Button>
    </div>
</main>

<style>
    main {
        display: flex;
        justify-content: center;
    }
    div.inner {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .mt {
        margin-top: 40px;
    }
</style>