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

export function useOnLoad(callback: () => void) {
  const firstRenderRef = useRef(true);
  useEffect(() => {
    if (!firstRenderRef.current) {
      return;
    }

    firstRenderRef.current = false;
    callback();
  }, [callback]);
}

export function usePrevious(value: any) {
  const ref = useRef();
  useEffect(() => {
    ref.current = value;
  });
  return ref.current;
}
