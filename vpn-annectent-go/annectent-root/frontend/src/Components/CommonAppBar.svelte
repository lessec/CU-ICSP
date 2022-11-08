<script>
    import TopAppBar, { Row, Section } from '@smui/top-app-bar';
    import IconButton from '@smui/icon-button';
    import LogoSmall from '../Assets/LogoSmall.svelte';
    import Drawer, {
        AppContent,
        Content,
        Header,
        Title,
        Subtitle,
        Scrim,
    } from '@smui/drawer';
    import List, { Item, Text } from '@smui/list';

    import { Auth } from 'aws-amplify';
    import { UserStore } from '../Stores/user-store.js';
    import { navigate } from "svelte-routing";


    export let isSetting;
    export let isRefresh;
    export let isDelete;
    export let action;

    let open = false;

    async function signOut() {
      try {
        await Auth.signOut().then(() => {
                  $UserStore = null;
                  navigate("/logout", {replace: true});
              });
      } catch (error) {
        console.log('error signing out: ', error);
      }
    }
</script>

<Drawer variant="modal" fixed={false} bind:open>
    <Header>
      <Title>Annectent</Title>
    </Header>
    <Content>
      <List>
        <Item
          on:click={() => {
            navigate("/admin", { replace: true });
            open = false;
          }}>
          <Text>Home</Text>
        </Item>
        <Item
          on:click={() => {
            navigate("/user", { replace: true });
            open = false;
          }}>
          <Text>User</Text>
        </Item>
        <Item
          on:click={() => {
            navigate("/usergroup", { replace: true });
            open = false;
          }}>
          <Text>Group</Text>
        </Item>
        <Item
            on:click={() => {
               navigate("/mfa", {replace: true})
               open = false; 
            }}>
            <Text>MFA</Text>
        </Item>
        <Item
            on:click={signOut}>
            <Text>Sign Out</Text>
        </Item>
      </List>
    </Content>
</Drawer>

<TopAppBar
      variant="static"
      color='secondary'>
      <Row>
        <Section>
          <IconButton class="material-icons" on:click={() => {open = !open}}>menu</IconButton>
        </Section>
        <Section align="center">
            <LogoSmall />
        </Section>
        <Section align="end" toolbar>
            {#if isSetting}
            <IconButton class="material-icons" aria-label="setting" on:click={action}>
                settings
            </IconButton>
            {/if}
            {#if isRefresh}
            <IconButton class="material-icons" aria-label="refresh" on:click={action}>
                loop
            </IconButton>
            {/if}
            {#if isDelete}
            <IconButton class="material-icons" aria-label="delete" on:click={action}>
                delete
            </IconButton>
            {/if}
        </Section>
      </Row>
</TopAppBar>