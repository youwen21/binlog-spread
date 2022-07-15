{{if $data._child}}
<li class="nav-item has-treeview">
    <a href="${pathname}" class="nav-link">
    <i class="nav-icon fas fa-tachometer-alt"></i>
    <p>
    ${title}
    <i class="right fas fa-angle-left"></i>
    </p>
    </a>
    <ul class="nav nav-treeview">
    {{tmpl($data._child) '#sidebar-level-1'}}
</ul>
</li>
{{else}}
<li class="nav-item">
    <a href="${pathname}" class="nav-link">
    <i class="nav-icon far fa-image"></i>
    <p>
    ${title}
    </p>
    </a>
    </li>
{{/if}}