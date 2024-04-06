/**
 * @see https://umijs.org/docs/max/access#access
 * */
export default function access(initialState: InitialState | undefined) {
  const { loginUser } = initialState ?? {};
  return {
    // userRole 为 0 时表示普通用户，为 1 时表示管理员
    canAdmin: loginUser?.userRole === 1,
  };
}
