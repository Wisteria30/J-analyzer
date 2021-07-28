export const ROUTES = {
  GET: {
    AUTH: '/api/v1/authorize',
    TOKEN: '/api/v1/token',
    LIST: '/api/v1/images',
    IMAGE: '/api/v1/images/:id',
  },
  POST: {
    ANALYZE: '/api/v1/images/:id',
  },
}

export default ROUTES
