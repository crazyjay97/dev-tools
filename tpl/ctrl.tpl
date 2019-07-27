package {{packageName}}.{{moduleName}}.controller;

import java.util.Arrays;
import java.util.Map;

import org.apache.shiro.authz.annotation.RequiresPermissions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import {{packageName}}.{{moduleName}}.entity.{{className}}Entity;
import {{packageName}}.{{moduleName}}.service.{{className}}Service;
import {{mainPath}}.common.utils.PageUtils;
import {{mainPath}}.common.utils.R;



/**
  * {{ table.TableComment }}
  *
  * @author {{ authorName }}
  * @email {{ emailAddress }}
  * @date {{ genTime }}
  */
@RestController
@RequestMapping("{{ moduleName }}/{{ fileName }}")
public class {{className}}Controller {
    @Autowired
    private {{className}}Service {{fileName}}Service;

    /**
      * 列表
      */
    @RequestMapping("/list")
    @RequiresPermissions("{{moduleName}}:{{fileName}}:list")
    public R list(@RequestParam Map<String, Object> params){
        PageUtils page = {{fileName}}Service.queryPage(params);
        return R.ok().put("page", page);
    }


    /**
      * 信息
      */
    @RequestMapping("/info/{{"{"}}{{pkColumn.FieldName}}{{"}"}}")
    @RequiresPermissions("{{moduleName}}:{{fileName}}:info")
    public R info(@PathVariable("{{pkColumn.FieldName}}") {{pkColumn.JavaType}} {{pkColumn.FieldName}}){
        {{className}}Entity {{fileName}} = {{FieldName}}Service.selectById({{pkColumn.FieldName}});
        return R.ok().put("{{fileName}}", {{fileName}});
    }

    /**
      * 保存
      */
    @RequestMapping("/save")
    @RequiresPermissions("{{moduleName}}:{{fileName}}:save")
    public R save(@RequestBody {{className}}Entity {{fileName}}){
        {{fileName}}Service.insert({{fileName}});
        return R.ok();
    }

    /**
      * 修改
      */
    @RequestMapping("/update")
    @RequiresPermissions("{{moduleName}}:{{fileName}}:update")
    public R update(@RequestBody {{className}}Entity {{fileName}}){
        {{fileName}}Service.update({{fileName}});
        return R.ok();
    }

    /**
      * 删除
      */
    @RequestMapping("/delete")
    @RequiresPermissions("{{moduleName}}:{{fileName}}:delete")
    public R delete(@RequestBody {{pkColumn.JavaType}}[] {{pkColumn.FieldName}}s){
        {{fileName}}Service.deleteBatchIds(Arrays.asList({{pkColumn.FieldName}}s));
        return R.ok();
    }

}
