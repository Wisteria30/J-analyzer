<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card width="400px" class="mx-auto mt-5">
        <v-card-title>
          <h1 class="display-1">Sign up</h1>
        </v-card-title>
        <v-card-text>
          <v-form>
            <v-text-field
              prepend-icon="mdi-account-circle"
              label="Email"
              v-model="email"
            />
            <v-text-field
              v-bind:type="isShowPassword ? 'text' : 'password'"
              v-bind:append-icon="isShowPassword ? 'mdi-eye' : 'mdi-eye-off'"
              @click:append="showPassword"
              v-model="password"
              prepend-icon="mdi-lock"
              label="Password"
            />
            <v-card-actions>
              <v-btn class="primary" @click="signUp">Sign up</v-btn>
            </v-card-actions>
          </v-form>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  computed: {
    isShowPassword() {
      return this.$store.getters.getSignUpShowPassword
    },
  },
  methods: {
    showPassword() {
      this.$store.dispatch('setSignUpShowPassword')
    },
    signUp() {
        const user = {
            email: this.email,
            password: this.password
        }
        this.$store.dispatch('signUp', user)
    },
  },
}
</script>
