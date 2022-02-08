;; Eval Buffer with `M-x eval-buffer' to register the newly created template.

(dap-register-debug-template
  "Launch Debug Food Delivery"
  (list :type "go"
        :request "launch"
        :name "Launch Debug Food Delivery"
        :mode "debug"
        :program "${workspaceFolder}/main.go"
        :buildFlags "-gcflags '-N -l'"
        :args nil
        :env nil
        :envFile nil))
