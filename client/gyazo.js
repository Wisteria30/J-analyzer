// import { addAuthorize, assignDefaults } from '@nuxtjs/auth-next/utils/provider'

export default function gyazo(
  // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/explicit-module-boundary-types
  nuxt,
  strategy
) {
  const DEFAULTS = {
    scheme: 'oauth2',
    name: 'gyazo',
    endpoints: {
      authorization: 'https://api.gyazo.com/oauth/authorize',
      token: 'https://api.gyazo.com/oauth/token',
    },
    token: {
      property: 'access_token',
      type: 'Bearer',
      maxAge: Infinity,
    },
    responseType: 'code',
    grantType: 'authorization_code',
  }

  // assignDefaults(strategy, DEFAULTS)

  // addAuthorize(nuxt, strategy, true)
}
