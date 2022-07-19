import { RefObject, useEffect, useRef } from 'react';

export type CheckboxProps = {
  checked: boolean;
  indeterminate: boolean;
};

export function useCheckbox({
  checked,
  indeterminate,
}: CheckboxProps): RefObject<HTMLInputElement> {
  const ref = useRef<HTMLInputElement>(null);
  useEffect(() => {
    if (!ref.current) {
      return;
    }

    ref.current.checked = checked;
    ref.current.indeterminate = indeterminate;
  }, [checked, indeterminate]);
  return ref;
}
