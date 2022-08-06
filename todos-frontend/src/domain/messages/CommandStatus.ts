export type CommandStatus = Readonly<{
  success: boolean;
  errorMessage?: string;
}>;

export function Success(): CommandStatus {
  return { success: true };
}

export function Failure(errorMessage: string): CommandStatus {
  return { success: false, errorMessage };
}
