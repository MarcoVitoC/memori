export const GetEnv = (key: string, fallback: string) => {
  const value = import.meta.env[key as keyof ImportMetaEnv]

  return (!value) ? fallback : value
}